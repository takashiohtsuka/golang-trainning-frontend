package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
)

type PrefectureUsecase interface {
	GetList(ctx context.Context) ([]querymodel.PrefectureQueryModel, error)
}
