package repository

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/querymodel"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
	"golang-trainning-frontend/pkg/usecase/outputport"
	"golang-trainning-frontend/pkg/usecase/query"

	"gorm.io/gorm"
)

type immediateAvailableWomanRepository struct {
	db *gorm.DB
}

func NewImmediateAvailableWomanRepository(db *gorm.DB) outputport.ImmediateAvailableWomanRepository {
	return &immediateAvailableWomanRepository{db: db}
}

const immediateAvailableWomanBaseSQL = `
	SELECT
		iaw.id          AS iaw_id,
		iaw.expires_at  AS iaw_expires_at,
		w.id            AS woman_id,
		w.name          AS woman_name,
		w.age,
		w.birthplace,
		w.blood_type,
		w.hobby,
		s.id            AS store_id,
		s.name          AS store_name,
		bt.code         AS business_type_code,
		wi.id           AS image_id,
		wi.path         AS image_path
	FROM immediate_available_women iaw
	JOIN women w          ON w.id = iaw.woman_id AND w.deleted_at IS NULL AND w.is_active = TRUE
	JOIN stores s         ON s.id = iaw.store_id AND s.deleted_at IS NULL AND s.is_active = TRUE
	JOIN districts d      ON d.id = s.district_id
	JOIN prefectures p    ON p.id = d.prefecture_id
	JOIN business_types bt ON bt.id = s.business_type_id
	LEFT JOIN woman_images wi ON wi.woman_id = w.id
	WHERE iaw.expires_at > NOW()`

func (r *immediateAvailableWomanRepository) FindAll(
	ctx context.Context,
	conditions []query.Condition,
	page uint,
	limit uint,
) (collection.Collection[querymodel.ImmediateAvailableWomanQueryModel], error) {
	offset := (page - 1) * limit

	where, args := buildWhereClause(conditions)
	sql := immediateAvailableWomanBaseSQL + where + " ORDER BY iaw.expires_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[querymodel.ImmediateAvailableWomanQueryModel](nil), err
	}

	return mapToImmediateAvailableWomanQueryModel(rows), nil
}

func (r *immediateAvailableWomanRepository) TotalCount(ctx context.Context, conditions []query.Condition) (uint, error) {
	where, args := buildWhereClause(conditions)
	countSQL := "SELECT COUNT(*) FROM (" + immediateAvailableWomanBaseSQL + where + ") AS sub"

	var total uint
	if err := r.db.WithContext(ctx).Raw(countSQL, args...).Scan(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func mapToImmediateAvailableWomanQueryModel(rows []map[string]any) collection.Collection[querymodel.ImmediateAvailableWomanQueryModel] {
	type key struct{ iawID uint }
	order := make([]key, 0)
	seen := make(map[key]bool)
	iawMap := make(map[key]*querymodel.ImmediateAvailableWoman)
	seenImages := make(map[key]map[uint]bool)

	for _, row := range rows {
		iawID := helper.ToUint(row["iaw_id"])
		k := key{iawID}

		if !seen[k] {
			seen[k] = true
			order = append(order, k)
			iawMap[k] = &querymodel.ImmediateAvailableWoman{
				ID:         helper.ToUint(row["woman_id"]),
				Name:       helper.ToString(row["woman_name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
				Store: querymodel.ImmediateAvailableWomanStore{
					ID:           helper.ToUint(row["store_id"]),
					Name:         helper.ToString(row["store_name"]),
					BusinessType: fvo.NewBusinessType(helper.ToString(row["business_type_code"])),
				},
				ExpiresAt: helper.ToTimePtr(row["iaw_expires_at"]),
			}
			seenImages[k] = make(map[uint]bool)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[k][imageID] {
			seenImages[k][imageID] = true
			current := iawMap[k].Images.All()
			current = append(current, querymodel.WomanImage{
				ID:   imageID,
				Path: helper.ToString(row["image_path"]),
			})
			iawMap[k].Images = collection.NewCollection(current)
		}
	}

	items := make([]querymodel.ImmediateAvailableWomanQueryModel, 0, len(order))
	for _, k := range order {
		items = append(items, iawMap[k])
	}
	return collection.NewCollection(items)
}
