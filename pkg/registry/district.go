package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewDistrictController() controller.District {
	u := interactor.NewDistrictUsecase(
		frontendrepository.NewDistrictRepository(r.db),
	)
	return controller.NewDistrictController(u)
}
