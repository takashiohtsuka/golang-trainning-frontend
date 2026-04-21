package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
)

type StoreUsecase interface {
	GetDetail(ctx context.Context, i input.GetStoreDetailInput) (querymodel.StoreQueryModel, error)
}
