package repository

import (
	"context"

	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type businessTypeRepository struct {
	db *gorm.DB
}

func NewBusinessTypeRepository(db *gorm.DB) outputport.BusinessTypeRepository {
	return &businessTypeRepository{db: db}
}

func (r *businessTypeRepository) FindAll(ctx context.Context) ([]querymodel.BusinessTypeQueryModel, error) {
	sql := `SELECT code FROM business_types ORDER BY id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.BusinessTypeQueryModel, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.BusinessType{
			Code: helper.ToString(row["code"]),
		})
	}
	return result, nil
}
