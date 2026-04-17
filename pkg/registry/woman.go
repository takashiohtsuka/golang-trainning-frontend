package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewWomanController() controller.Woman {
	u := interactor.NewWomanUsecase(
		frontendrepository.NewWomanRepository(r.db),
	)
	return controller.NewWomanController(u)
}
