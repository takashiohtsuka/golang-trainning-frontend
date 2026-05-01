package controller

import (
	"net/http"

	responseBusinessTypes "golang-trainning-frontend/pkg/adapter/response/business_types"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type businessTypeController struct {
	usecase inputport.BusinessTypeUsecase
}

type BusinessType interface {
	GetBusinessTypeList(c Context) error
}

func NewBusinessTypeController(u inputport.BusinessTypeUsecase) BusinessType {
	return &businessTypeController{usecase: u}
}

func (bc *businessTypeController) GetBusinessTypeList(ctx Context) error {
	businessTypes, err := bc.usecase.GetList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseBusinessTypes.NewListResponse(businessTypes))
}
