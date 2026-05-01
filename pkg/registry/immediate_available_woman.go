package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewImmediateAvailableWomanController() controller.ImmediateAvailableWoman {
	u := interactor.NewImmediateAvailableWomanUsecase(
		frontendrepository.NewImmediateAvailableWomanRepository(r.db),
	)
	return controller.NewImmediateAvailableWomanController(u)
}
