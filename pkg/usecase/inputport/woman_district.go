package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanDistrictUsecase interface {
	GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[dto.WomanDTO], uint, error)
}
