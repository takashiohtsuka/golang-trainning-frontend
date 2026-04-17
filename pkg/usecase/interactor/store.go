package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
	"golang-trainning-frontend/pkg/usecase/query"
)

type storeUsecase struct {
	storeRepository outputport.StoreRepository
}

func NewStoreUsecase(storeRepository outputport.StoreRepository) inputport.StoreUsecase {
	return &storeUsecase{storeRepository}
}

func (u *storeUsecase) GetDetail(ctx context.Context, i input.GetStoreDetailInput) (entity.StoreEntity, error) {
	return u.storeRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.StoreID),
	})
}
