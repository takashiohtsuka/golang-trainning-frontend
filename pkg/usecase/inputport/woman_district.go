package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanDistrictUsecase interface {
	GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[entity.WomanEntity], error)
}
