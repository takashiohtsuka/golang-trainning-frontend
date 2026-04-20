package controller

import (
	"net/http"
	"strconv"

	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/input"
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
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	query := ctx.Request().URL.Query()

	count, err := wc.womanDistrictCountUsecase.GetCount(ctx.Request().Context(), input.GetWomanDistrictCountInput{
		DistrictID: uint(id),
		BloodTypes: query["blood_type"],
		AgeRanges:  query["age_range"],
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDistrictCountResponse(count))
}
