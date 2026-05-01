package registry

import (
	"golang-trainning-frontend/pkg/adapter/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Store:                   r.NewStoreController(),
		Woman:                   r.NewWomanController(),
		WomanDistrict:           r.NewWomanDistrictController(),
		WomanDistrictCount:      r.NewWomanDistrictCountController(),
		ImmediateAvailableWoman: r.NewImmediateAvailableWomanController(),
		Prefecture:              r.NewPrefectureController(),
		District:                r.NewDistrictController(),
		BusinessType:            r.NewBusinessTypeController(),
	}
}
