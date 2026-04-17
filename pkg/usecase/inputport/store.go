package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/input"
)

type StoreUsecase interface {
	GetDetail(ctx context.Context, i input.GetStoreDetailInput) (entity.StoreEntity, error)
}
