package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/query"
)

type StoreDistrictRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.StoreEntity], error)
}
