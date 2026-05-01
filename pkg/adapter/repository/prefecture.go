package repository

import (
	"context"

	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type prefectureRepository struct {
	db *gorm.DB
}

func NewPrefectureRepository(db *gorm.DB) outputport.PrefectureRepository {
	return &prefectureRepository{db: db}
}

func (r *prefectureRepository) FindAll(ctx context.Context) ([]querymodel.PrefectureQueryModel, error) {
	sql := `SELECT id, name FROM prefectures ORDER BY id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.PrefectureQueryModel, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.Prefecture{
			ID:   helper.ToUint(row["id"]),
			Name: helper.ToString(row["name"]),
		})
	}
	return result, nil
}
