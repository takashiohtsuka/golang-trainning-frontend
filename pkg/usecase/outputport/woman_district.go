package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanDistrictRepository interface {
	FindAllByDistrict(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[querymodel.WomanQueryModel], error)
	CountByDistrictWithCondition(ctx context.Context, i input.GetWomanDistrictCountInput) (uint, error)
}
