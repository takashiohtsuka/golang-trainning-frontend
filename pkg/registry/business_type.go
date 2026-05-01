package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewBusinessTypeController() controller.BusinessType {
	u := interactor.NewBusinessTypeUsecase(
		frontendrepository.NewBusinessTypeRepository(r.db),
	)
	return controller.NewBusinessTypeController(u)
}
