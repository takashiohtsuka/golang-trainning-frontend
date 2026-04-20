package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/dto"
	"golang-trainning-frontend/pkg/usecase/query"
)

type StoreRepository interface {
	FindOne(ctx context.Context, conditions []query.Condition) (dto.StoreDTO, error)
}
