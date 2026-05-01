package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type DistrictUsecase interface {
	GetListByPrefecture(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQueryModel, error)
}
