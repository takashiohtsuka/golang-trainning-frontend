package interactor

import (
	"context"

	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/inputport"
	"golang-trainning-frontend/pkg/usecase/outputport"
)

type businessTypeUsecase struct {
	repo outputport.BusinessTypeRepository
}

func NewBusinessTypeUsecase(repo outputport.BusinessTypeRepository) inputport.BusinessTypeUsecase {
	return &businessTypeUsecase{repo: repo}
}

func (u *businessTypeUsecase) GetList(ctx context.Context) ([]querymodel.BusinessTypeQueryModel, error) {
	return u.repo.FindAll(ctx)
}
