package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
	"golang-trainning-frontend/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[dto.WomanDTO], error)
	FindOne(ctx context.Context, conditions []query.Condition) (dto.WomanDTO, error)
}
