package interactor

import (
	"context"
	"sync"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/dto"
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

func (u *womanDistrictUsecase) GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[dto.WomanDTO], uint, error) {
	if i.Page == 0 {
		i.Page = 1
	}

	var (
		women collection.Collection[dto.WomanDTO]
		total uint
		err1  error
		err2  error
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		women, err1 = u.womanDistrictRepository.FindAllByDistrict(ctx, i)
	}()

	go func() {
		defer wg.Done()
		total, err2 = u.womanDistrictRepository.CountByDistrictWithCondition(ctx, input.GetWomanDistrictCountInput{
			DistrictID: i.DistrictID,
			BloodTypes: i.BloodTypes,
			AgeRanges:  i.AgeRanges,
		})
	}()

	wg.Wait()

	if err1 != nil {
		return collection.NewCollection[dto.WomanDTO](nil), 0, err1
	}
	if err2 != nil {
		return collection.NewCollection[dto.WomanDTO](nil), 0, err2
	}

	return women, total, nil
}
