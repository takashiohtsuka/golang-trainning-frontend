package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanDistrictCountUsecase interface {
	GetCount(ctx context.Context, i input.GetWomanDistrictCountInput) (uint, error)
}
