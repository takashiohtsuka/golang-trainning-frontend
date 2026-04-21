package store

import (
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
	"golang-trainning-frontend/pkg/helper"
)

// MapToQueryModel は flat な rows を StoreQueryModel のコレクションに変換する。
func MapToQueryModel(rows []map[string]any) collection.Collection[querymodel.StoreQueryModel] {
	storeOrder := make([]uint, 0)
	storeMap := make(map[uint]*querymodel.Store)
	womanOrderByStore := make(map[uint][]uint)
	womanMap := make(map[uint]*querymodel.Woman)
	seenWomenByStore := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		storeID := helper.ToUint(row["store_id"])

		if _, exists := storeMap[storeID]; !exists {
			storeOrder = append(storeOrder, storeID)
			storeMap[storeID] = &querymodel.Store{
				ID:           storeID,
				District:     fvo.NewDistrict(helper.ToUint(row["district_id"]), helper.ToString(row["district_name"])),
				Prefecture:   fvo.NewPrefecture(helper.ToUint(row["prefecture_id"]), helper.ToString(row["prefecture_name"])),
				Region:       fvo.NewRegion(helper.ToUint(row["region_id"]), helper.ToString(row["region_name"])),
				BusinessType: fvo.NewBusinessType(helper.ToString(row["business_type_code"])),
				Name:         helper.ToString(row["store_name"]),
			}
			womanOrderByStore[storeID] = make([]uint, 0)
		}

		womanID := helper.ToUint(row["woman_id"])
		if womanID == 0 {
			continue
		}

		if seenWomenByStore[storeID] == nil {
			seenWomenByStore[storeID] = make(map[uint]bool)
		}
		if !seenWomenByStore[storeID][womanID] {
			seenWomenByStore[storeID][womanID] = true
			womanOrderByStore[storeID] = append(womanOrderByStore[storeID], womanID)
		}

		if _, exists := womanMap[womanID]; !exists {
			womanMap[womanID] = &querymodel.Woman{
				ID:         womanID,
				Name:       helper.ToString(row["woman_name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
			}
			seenBlogs[womanID] = make(map[uint]bool)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[womanID][blogID] {
			seenBlogs[womanID][blogID] = true
			current := womanMap[womanID].Blogs.All()
			current = append(current, &querymodel.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       helper.ToString(row["blog_title"]),
				IsPublished: true,
				Photos:      collection.NewCollection[querymodel.Photo](nil),
			})
			womanMap[womanID].Blogs = collection.NewCollection(current)
		}
	}

	for storeID, womanIDs := range womanOrderByStore {
		women := make([]querymodel.WomanQueryModel, 0, len(womanIDs))
		for _, wid := range womanIDs {
			women = append(women, womanMap[wid])
		}
		storeMap[storeID].Women = collection.NewCollection(women)
	}

	items := make([]querymodel.StoreQueryModel, 0, len(storeOrder))
	for _, storeID := range storeOrder {
		items = append(items, storeMap[storeID])
	}
	return collection.NewCollection(items)
}
