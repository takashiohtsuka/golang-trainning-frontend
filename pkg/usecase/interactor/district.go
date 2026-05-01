package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
)

type districtUsecase struct {
	repo outputport.DistrictRepository
}

func NewDistrictUsecase(repo outputport.DistrictRepository) inputport.DistrictUsecase {
	return &districtUsecase{repo: repo}
}

func (u *districtUsecase) GetListByPrefecture(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQueryModel, error) {
	return u.repo.FindAllByPrefecture(ctx, prefectureID)
}
