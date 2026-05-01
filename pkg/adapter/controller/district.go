package controller

import (
	"net/http"

	responseDistricts "golang-trainning-frontend/pkg/adapter/response/districts"
	requestPrefectures "golang-trainning-frontend/pkg/adapter/request/prefectures"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type districtController struct {
	usecase inputport.DistrictUsecase
}

type District interface {
	GetDistrictList(c Context) error
}

func NewDistrictController(u inputport.DistrictUsecase) District {
	return &districtController{usecase: u}
}

func (dc *districtController) GetDistrictList(ctx Context) error {
	var req requestPrefectures.DistrictListRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	districts, err := dc.usecase.GetListByPrefecture(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseDistricts.NewListResponse(districts))
}
