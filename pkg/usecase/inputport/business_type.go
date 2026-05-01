package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type BusinessTypeUsecase interface {
	GetList(ctx context.Context) ([]querymodel.BusinessTypeQueryModel, error)
}
