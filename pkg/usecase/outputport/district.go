package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type DistrictRepository interface {
	FindAllByPrefecture(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQueryModel, error)
}
