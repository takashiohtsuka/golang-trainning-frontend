package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewWomanDistrictCountController() controller.WomanDistrictCount {
	u := interactor.NewWomanDistrictCountUsecase(
		frontendrepository.NewWomanDistrictRepository(r.db),
	)
	return controller.NewWomanDistrictCountController(u)
}
