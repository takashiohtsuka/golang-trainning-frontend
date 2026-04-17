package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.WomanEntity, error)
}
