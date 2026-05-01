package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
)

type prefectureUsecase struct {
	repo outputport.PrefectureRepository
}

func NewPrefectureUsecase(repo outputport.PrefectureRepository) inputport.PrefectureUsecase {
	return &prefectureUsecase{repo: repo}
}

func (u *prefectureUsecase) GetList(ctx context.Context) ([]querymodel.PrefectureQueryModel, error) {
	return u.repo.FindAll(ctx)
}
