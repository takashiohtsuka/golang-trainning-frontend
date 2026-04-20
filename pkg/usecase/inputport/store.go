package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/dto"
	"golang-trainning-frontend/pkg/usecase/input"
)

type StoreUsecase interface {
	GetDetail(ctx context.Context, i input.GetStoreDetailInput) (dto.StoreDTO, error)
}
