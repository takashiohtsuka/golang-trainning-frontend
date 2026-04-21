package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/apperror"
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
	"golang-trainning-frontend/pkg/usecase/query"
)

type womanUsecase struct {
	womanRepository outputport.WomanRepository
}

func NewWomanUsecase(womanRepository outputport.WomanRepository) inputport.WomanUsecase {
	return &womanUsecase{womanRepository}
}

func (u *womanUsecase) GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error) {
	return u.womanRepository.FindAll(ctx, []query.Condition{})
}

func (u *womanUsecase) GetStoreWomanList(ctx context.Context, i input.GetStoreWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error) {
	return u.womanRepository.FindAll(ctx, []query.Condition{
		query.Where("wsa.store_id", i.StoreID),
	})
}

func (u *womanUsecase) GetDetail(ctx context.Context, i input.GetWomanDetailInput) (querymodel.WomanQueryModel, error) {
	woman, err := u.womanRepository.FindOne(ctx, []query.Condition{
		query.Where("w.id", i.WomanID),
	})
	if err != nil {
		return nil, err
	}
	if woman.IsNil() {
		return nil, apperror.NewNotFoundException("woman not found")
	}
	return woman, nil
}
