package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[querymodel.WomanQueryModel], error)
	FindOne(ctx context.Context, conditions []query.Condition) (querymodel.WomanQueryModel, error)
}
