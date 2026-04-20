package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
)

type womanDistrictCountUsecase struct {
	womanDistrictRepository outputport.WomanDistrictRepository
}

func NewWomanDistrictCountUsecase(r outputport.WomanDistrictRepository) inputport.WomanDistrictCountUsecase {
	return &womanDistrictCountUsecase{r}
}

func (u *womanDistrictCountUsecase) GetCount(ctx context.Context, i input.GetWomanDistrictCountInput) (uint, error) {
	return u.womanDistrictRepository.CountByDistrictWithCondition(ctx, i)
}
