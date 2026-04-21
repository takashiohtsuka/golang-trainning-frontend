package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/query"
)

type StoreDistrictRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[querymodel.StoreQueryModel], error)
}
