package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
)

type ImmediateAvailableWomanUsecase interface {
	GetList(ctx context.Context, i input.GetImmediateAvailableWomanListInput) (collection.Collection[querymodel.ImmediateAvailableWomanQueryModel], uint, error)
}
