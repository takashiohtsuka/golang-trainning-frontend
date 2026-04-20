package woman

import (
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
	fvo "golang-trainning-frontend/pkg/dto/valueobject"
	"golang-trainning-frontend/pkg/helper"
)

// MapToDTO は flat な rows を Woman DTO のコレクションに変換する。
func MapToDTO(rows []map[string]any) collection.Collection[dto.WomanDTO] {
	womanOrder := make([]uint, 0)
	womanMap := make(map[uint]*dto.Woman)
	seenStores := make(map[uint]map[uint]bool)
	seenImages := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		womanID := helper.ToUint(row["woman_id"])

		if _, exists := womanMap[womanID]; !exists {
			womanOrder = append(womanOrder, womanID)
			womanMap[womanID] = &dto.Woman{
				ID:         womanID,
				Name:       helper.ToString(row["woman_name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
			}
			seenStores[womanID] = make(map[uint]bool)
			seenImages[womanID] = make(map[uint]bool)
			seenBlogs[womanID] = make(map[uint]bool)
		}

		assignmentID := helper.ToUint(row["assignment_id"])
		if assignmentID != 0 && !seenStores[womanID][assignmentID] {
			seenStores[womanID][assignmentID] = true
			stores := womanMap[womanID].Stores.All()
			stores = append(stores, dto.WomanStore{
				ID:           helper.ToUint(row["assignment_store_id"]),
				Name:         helper.ToString(row["assignment_store_name"]),
				BusinessType: fvo.NewBusinessType(helper.ToString(row["assignment_store_business_type"])),
			})
			womanMap[womanID].Stores = collection.NewCollection(stores)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[womanID][imageID] {
			seenImages[womanID][imageID] = true
			current := womanMap[womanID].Images.All()
			current = append(current, dto.WomanImage{
				ID:   imageID,
				Path: helper.ToString(row["image_path"]),
			})
			womanMap[womanID].Images = collection.NewCollection(current)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[womanID][blogID] {
			seenBlogs[womanID][blogID] = true
			current := womanMap[womanID].Blogs.All()
			current = append(current, &dto.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       helper.ToString(row["blog_title"]),
				IsPublished: true,
				Photos:      collection.NewCollection[dto.Photo](nil),
			})
			womanMap[womanID].Blogs = collection.NewCollection(current)
		}
	}

	items := make([]dto.WomanDTO, 0, len(womanOrder))
	for _, wid := range womanOrder {
		items = append(items, womanMap[wid])
	}
	return collection.NewCollection(items)
}
