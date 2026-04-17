package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewStoreController() controller.Store {
	u := interactor.NewStoreUsecase(
		frontendrepository.NewStoreRepository(r.db),
	)
	return controller.NewStoreController(u)
}
