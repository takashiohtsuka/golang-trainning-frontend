package repository

import (
	"context"

	womanMapper "golang-trainning-frontend/pkg/adapter/mapper/woman"
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

const limit uint = 10
const womanDistrictBlogsLimit = 3

type womanDistrictRepository struct {
	db *gorm.DB
}

func NewWomanDistrictRepository(db *gorm.DB) outputport.WomanDistrictRepository {
	return &womanDistrictRepository{db: db}
}

func (r *womanDistrictRepository) CountByDistrictWithCondition(ctx context.Context, i input.GetWomanDistrictCountInput) (uint, error) {
	sql := `
		SELECT COUNT(DISTINCT w.id, wsa.id)
		FROM women w
		JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		JOIN stores s ON s.id = wsa.store_id AND s.deleted_at IS NULL AND s.is_active = TRUE
		JOIN districts d ON s.district_id = d.id
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE
		AND d.id = ?`

	args := []any{i.DistrictID}
	condition, filterArgs := buildWomanFilterCondition(i.BloodTypes, i.AgeRanges)
	sql += condition
	args = append(args, filterArgs...)

	var total uint
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *womanDistrictRepository) FindAllByDistrict(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[querymodel.WomanQueryModel], error) {
	offset := (i.Page - 1) * limit

	subQuery := `
		SELECT DISTINCT w.id AS woman_id, wsa.id AS assignment_id
		FROM women w
		JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		JOIN stores s   ON s.id = wsa.store_id AND s.deleted_at IS NULL AND s.is_active = TRUE
		JOIN districts d ON s.district_id = d.id
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE
		AND d.id = ?`

	args := []any{i.DistrictID}
	condition, filterArgs := buildWomanFilterCondition(i.BloodTypes, i.AgeRanges)
	subQuery += condition
	args = append(args, filterArgs...)

	subQuery += " ORDER BY w.id, wsa.id LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	sql := `
		SELECT
			w.id            AS woman_id,
			w.name          AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			wsa.id          AS assignment_id,
			wsa.store_id    AS assignment_store_id,
			s.name          AS assignment_store_name,
			bt.code         AS assignment_store_business_type,
			wi.id           AS image_id,
			wi.path         AS image_path,
			b.id            AS blog_id,
			b.title         AS blog_title
		FROM (` + subQuery + `) AS paged
		JOIN women w ON w.id = paged.woman_id
		JOIN woman_store_assignments wsa ON wsa.id = paged.assignment_id
		JOIN stores s ON s.id = wsa.store_id AND s.deleted_at IS NULL
		LEFT JOIN business_types bt ON bt.id = s.business_type_id
		LEFT JOIN woman_images wi ON wi.woman_id = w.id
		LEFT JOIN (
			SELECT id, woman_id, title,
			       ROW_NUMBER() OVER (PARTITION BY woman_id ORDER BY created_at DESC) AS rn
			FROM blogs
			WHERE deleted_at IS NULL AND is_published = TRUE
		) b ON b.woman_id = w.id AND b.rn <= ?
		ORDER BY w.id, wsa.id, wi.id, b.id`

	args = append(args, womanDistrictBlogsLimit)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[querymodel.WomanQueryModel](nil), err
	}
	return womanMapper.MapToQueryModel(rows), nil
}
