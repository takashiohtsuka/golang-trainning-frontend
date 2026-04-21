package controller

import (
	"net/http"

	requestDistricts "golang-trainning-frontend/pkg/adapter/request/districts"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type womanDistrictCountController struct {
	womanDistrictCountUsecase inputport.WomanDistrictCountUsecase
}

type WomanDistrictCount interface {
	GetWomanDistrictCount(c Context) error
}

func NewWomanDistrictCountController(u inputport.WomanDistrictCountUsecase) WomanDistrictCount {
	return &womanDistrictCountController{u}
}

func (wc *womanDistrictCountController) GetWomanDistrictCount(ctx Context) error {
	var req requestDistricts.WomanCountRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	count, err := wc.womanDistrictCountUsecase.GetCount(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDistrictCountResponse(count))
}
