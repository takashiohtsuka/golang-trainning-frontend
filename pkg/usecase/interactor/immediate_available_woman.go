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

type immediateAvailableWomanUsecase struct {
	repo outputport.ImmediateAvailableWomanRepository
}

func NewImmediateAvailableWomanUsecase(repo outputport.ImmediateAvailableWomanRepository) inputport.ImmediateAvailableWomanUsecase {
	return &immediateAvailableWomanUsecase{repo: repo}
}

func (u *immediateAvailableWomanUsecase) GetList(ctx context.Context, i input.GetImmediateAvailableWomanListInput) (collection.Collection[querymodel.ImmediateAvailableWomanQueryModel], uint, error) {
	if i.Page == 0 {
		i.Page = 1
	}
	if i.Limit == 0 {
		i.Limit = 10
	}

	conditions := buildImmediateAvailableWomanConditions(i)

	var (
		women collection.Collection[querymodel.ImmediateAvailableWomanQueryModel]
		total uint
		err1  error
		err2  error
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		women, err1 = u.repo.FindAll(ctx, conditions, i.Page, i.Limit)
	}()

	go func() {
		defer wg.Done()
		total, err2 = u.repo.TotalCount(ctx, conditions)
	}()

	wg.Wait()

	if err1 != nil {
		return collection.NewCollection[querymodel.ImmediateAvailableWomanQueryModel](nil), 0, err1
	}
	if err2 != nil {
		return collection.NewCollection[querymodel.ImmediateAvailableWomanQueryModel](nil), 0, err2
	}

	return women, total, nil
}

func buildImmediateAvailableWomanConditions(i input.GetImmediateAvailableWomanListInput) []query.Condition {
	var conditions []query.Condition

	if i.PrefectureID != 0 {
		conditions = append(conditions, query.Where("p.id", i.PrefectureID))
	}
	if i.DistrictID != 0 {
		conditions = append(conditions, query.Where("d.id", i.DistrictID))
	}
	if len(i.BusinessTypes) > 0 {
		conditions = append(conditions, query.WhereIn("bt.code", i.BusinessTypes))
	}
	if len(i.BloodTypes) > 0 {
		conditions = append(conditions, query.WhereIn("w.blood_type", i.BloodTypes))
	}
	for _, ar := range i.AgeRanges {
		parts := strings.Split(ar, "-")
		if len(parts) == 2 {
			conditions = append(conditions, query.WhereBetweenOr("w.age", parts[0], parts[1]))
		}
	}

	return conditions
}
