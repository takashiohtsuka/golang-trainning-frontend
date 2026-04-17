package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/query"
)

type StoreRepository interface {
	FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error)
}
