package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/query"
)

type ImmediateAvailableWomanRepository interface {
	FindAll(
		ctx context.Context,
		conditions []query.Condition,
		page uint,
		limit uint,
	) (collection.Collection[querymodel.ImmediateAvailableWomanQueryModel], error)
	TotalCount(ctx context.Context, conditions []query.Condition) (uint, error)
}
