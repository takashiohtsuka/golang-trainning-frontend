package outputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type PrefectureRepository interface {
	FindAll(ctx context.Context) ([]querymodel.PrefectureQueryModel, error)
}
