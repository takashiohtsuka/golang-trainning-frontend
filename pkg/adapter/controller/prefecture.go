package controller

import (
	"net/http"

	responsePrefectures "golang-trainning-frontend/pkg/adapter/response/prefectures"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type prefectureController struct {
	usecase inputport.PrefectureUsecase
}

type Prefecture interface {
	GetPrefectureList(c Context) error
}

func NewPrefectureController(u inputport.PrefectureUsecase) Prefecture {
	return &prefectureController{usecase: u}
}

func (pc *prefectureController) GetPrefectureList(ctx Context) error {
	prefectures, err := pc.usecase.GetList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responsePrefectures.NewListResponse(prefectures))
}
