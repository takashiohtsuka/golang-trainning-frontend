package controller

import (
	"net/http"

	requestDistricts "golang-trainning-frontend/pkg/adapter/request/districts"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type womanDistrictController struct {
	womanDistrictUsecase inputport.WomanDistrictUsecase
}

type WomanDistrict interface {
	GetWomanDistrictList(c Context) error
}

func NewWomanDistrictController(u inputport.WomanDistrictUsecase) WomanDistrict {
	return &womanDistrictController{u}
}

func (wc *womanDistrictController) GetWomanDistrictList(ctx Context) error {
	var req requestDistricts.WomanListRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	women, total, err := wc.womanDistrictUsecase.GetList(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDistrictListResponse(women.All(), total))
}
