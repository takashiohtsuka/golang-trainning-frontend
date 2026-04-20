package controller

import (
	"net/http"
	"strconv"

	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/input"
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
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	page, _ := strconv.ParseUint(ctx.QueryParam("page"), 10, 64)
	query := ctx.Request().URL.Query()

	women, total, err := wc.womanDistrictUsecase.GetList(ctx.Request().Context(), input.GetWomanDistrictListInput{
		DistrictID: uint(id),
		Page:       uint(page),
		BloodTypes: query["blood_type"],
		AgeRanges:  query["age_range"],
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDistrictListResponse(women.All(), total))
}
