package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
)

type womanDistrictUsecase struct {
	womanDistrictRepository outputport.WomanDistrictRepository
}

func NewWomanDistrictUsecase(womanDistrictRepository outputport.WomanDistrictRepository) inputport.WomanDistrictUsecase {
	return &womanDistrictUsecase{womanDistrictRepository}
}

func (u *womanDistrictUsecase) GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[entity.WomanEntity], error) {
	return u.womanDistrictRepository.FindAllByDistrict(ctx, i.DistrictID)
}
