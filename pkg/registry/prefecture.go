package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	frontendrepository "golang-trainning-frontend/pkg/adapter/repository"
	"golang-trainning-frontend/pkg/usecase/interactor"
)

func (r *registry) NewPrefectureController() controller.Prefecture {
	u := interactor.NewPrefectureUsecase(
		frontendrepository.NewPrefectureRepository(r.db),
	)
	return controller.NewPrefectureController(u)
}
