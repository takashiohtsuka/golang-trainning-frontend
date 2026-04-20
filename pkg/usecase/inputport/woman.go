package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanUsecase interface {
	GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[dto.WomanDTO], error)
	GetStoreWomanList(ctx context.Context, i input.GetStoreWomanListInput) (collection.Collection[dto.WomanDTO], error)
	GetDetail(ctx context.Context, i input.GetWomanDetailInput) (dto.WomanDTO, error)
}
