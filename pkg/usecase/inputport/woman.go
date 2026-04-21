package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanUsecase interface {
	GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error)
	GetStoreWomanList(ctx context.Context, i input.GetStoreWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error)
	GetDetail(ctx context.Context, i input.GetWomanDetailInput) (querymodel.WomanQueryModel, error)
}
