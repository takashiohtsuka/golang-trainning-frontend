package woman

import (
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
	"golang-trainning-frontend/pkg/helper"
)

// MapToQueryModel は flat な rows を WomanQueryModel のコレクションに変換する。
// assignment_id をグルーピングキーとし、1エントリ = 1女性×1店舗 として扱う。
func MapToQueryModel(rows []map[string]any) collection.Collection[querymodel.WomanQueryModel] {
	assignmentOrder := make([]uint, 0)
	assignmentMap := make(map[uint]*querymodel.Woman)
	seenImages := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		assignmentID := helper.ToUint(row["assignment_id"])
		womanID := helper.ToUint(row["woman_id"])

		if _, exists := assignmentMap[assignmentID]; !exists {
			assignmentOrder = append(assignmentOrder, assignmentID)
			assignmentMap[assignmentID] = &querymodel.Woman{
				ID:         womanID,
				Name:       helper.ToString(row["woman_name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
				Stores: collection.NewCollection([]querymodel.WomanStore{
					{
						ID:           helper.ToUint(row["assignment_store_id"]),
						Name:         helper.ToString(row["assignment_store_name"]),
						BusinessType: fvo.NewBusinessType(helper.ToString(row["assignment_store_business_type"])),
					},
				}),
			}
			seenImages[assignmentID] = make(map[uint]bool)
			seenBlogs[assignmentID] = make(map[uint]bool)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[assignmentID][imageID] {
			seenImages[assignmentID][imageID] = true
			current := assignmentMap[assignmentID].Images.All()
			current = append(current, querymodel.WomanImage{
				ID:   imageID,
				Path: helper.ToString(row["image_path"]),
			})
			assignmentMap[assignmentID].Images = collection.NewCollection(current)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[assignmentID][blogID] {
			seenBlogs[assignmentID][blogID] = true
			current := assignmentMap[assignmentID].Blogs.All()
			current = append(current, &querymodel.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       helper.ToString(row["blog_title"]),
				IsPublished: true,
				Photos:      collection.NewCollection[querymodel.Photo](nil),
			})
			assignmentMap[assignmentID].Blogs = collection.NewCollection(current)
		}
	}

	items := make([]querymodel.WomanQueryModel, 0, len(assignmentOrder))
	for _, aid := range assignmentOrder {
		items = append(items, assignmentMap[aid])
	}
	return collection.NewCollection(items)
}
