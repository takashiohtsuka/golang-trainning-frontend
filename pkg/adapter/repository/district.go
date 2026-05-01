package repository

import (
	"context"

	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type districtRepository struct {
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) outputport.DistrictRepository {
	return &districtRepository{db: db}
}

func (r *districtRepository) FindAllByPrefecture(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQueryModel, error) {
	sql := `SELECT id, name, prefecture_id FROM districts WHERE prefecture_id = ? ORDER BY id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, prefectureID).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.DistrictQueryModel, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.District{
			ID:           helper.ToUint(row["id"]),
			Name:         helper.ToString(row["name"]),
			PrefectureID: helper.ToUint(row["prefecture_id"]),
		})
	}
	return result, nil
}
