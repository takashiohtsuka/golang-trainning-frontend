package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/query"
)

type StoreRepository interface {
	FindOne(ctx context.Context, conditions []query.Condition) (querymodel.StoreQueryModel, error)
}
