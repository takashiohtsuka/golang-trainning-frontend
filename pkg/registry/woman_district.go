package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewWomanDistrictController() controller.WomanDistrict {
	u := interactor.NewWomanDistrictUsecase(
		frontendrepository.NewWomanDistrictRepository(r.db),
	)
	return controller.NewWomanDistrictController(u)
}
