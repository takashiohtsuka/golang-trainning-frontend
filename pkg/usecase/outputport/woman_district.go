package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/query"
)

type WomanDistrictRepository interface {
	FindAllByDistrict(ctx context.Context, conditions []query.Condition, page uint) (collection.Collection[querymodel.WomanQueryModel], error)
	CountByDistrict(ctx context.Context, conditions []query.Condition) (uint, error)
}
