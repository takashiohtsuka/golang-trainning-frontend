package interactor

import (
	"context"
	"strings"
	"sync"

	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
	"golang-trainning-frontend/pkg/usecase/query"
)

type womanDistrictUsecase struct {
	womanDistrictRepository outputport.WomanDistrictRepository
}

func NewWomanDistrictUsecase(womanDistrictRepository outputport.WomanDistrictRepository) inputport.WomanDistrictUsecase {
	return &womanDistrictUsecase{womanDistrictRepository}
}

func (u *womanDistrictUsecase) GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[querymodel.WomanQueryModel], uint, error) {
	if i.Page == 0 {
		i.Page = 1
	}

	conditions := buildWomanDistrictConditions(i.DistrictID, i.BloodTypes, i.AgeRanges)

	var (
		women collection.Collection[querymodel.WomanQueryModel]
		total uint
		err1  error
		err2  error
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		women, err1 = u.womanDistrictRepository.FindAllByDistrict(ctx, conditions, i.Page)
	}()

	go func() {
		defer wg.Done()
		total, err2 = u.womanDistrictRepository.CountByDistrict(ctx, conditions)
	}()

	wg.Wait()

	if err1 != nil {
		return collection.NewCollection[querymodel.WomanQueryModel](nil), 0, err1
	}
	if err2 != nil {
		return collection.NewCollection[querymodel.WomanQueryModel](nil), 0, err2
	}

	return women, total, nil
}

// buildWomanDistrictConditions は district ID・血液型・年齢帯から conditions を組み立てる。
func buildWomanDistrictConditions(districtID uint, bloodTypes []string, ageRanges []string) []query.Condition {
	conditions := []query.Condition{
		query.Where("d.id", districtID),
	}
	if len(bloodTypes) > 0 {
		conditions = append(conditions, query.WhereIn("w.blood_type", bloodTypes))
	}
	for _, ar := range ageRanges {
		parts := strings.Split(ar, "-")
		if len(parts) == 2 {
			conditions = append(conditions, query.WhereBetweenOr("w.age", parts[0], parts[1]))
		}
	}
	return conditions
}
