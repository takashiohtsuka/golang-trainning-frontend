package controller

import (
	"errors"
	"net/http"

	"golang-trainning-frontend/pkg/apperror"
	requestStores "golang-trainning-frontend/pkg/adapter/request/stores"
	requestWomen "golang-trainning-frontend/pkg/adapter/request/women"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type womanController struct {
	womanUsecase inputport.WomanUsecase
}

type Woman interface {
	GetWomanList(c Context) error
	GetStoreWomanList(c Context) error
	GetWomanDetail(c Context) error
}

func NewWomanController(u inputport.WomanUsecase) Woman {
	return &womanController{u}
}

func (wc *womanController) GetWomanList(ctx Context) error {
	women, err := wc.womanUsecase.GetList(ctx.Request().Context(), input.GetWomanListInput{})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return ctx.JSON(http.StatusOK, responseWomen.NewListResponse(women.All()))
}

func (wc *womanController) GetStoreWomanList(ctx Context) error {
	var req requestStores.GetRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	women, err := wc.womanUsecase.GetStoreWomanList(ctx.Request().Context(), req.ToStoreWomanListInput())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	return ctx.JSON(http.StatusOK, responseWomen.NewListResponse(women.All()))
}

func (wc *womanController) GetWomanDetail(ctx Context) error {
	var req requestWomen.DetailRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	woman, err := wc.womanUsecase.GetDetail(ctx.Request().Context(), req.ToInput())
	if err != nil {
		var nfe *apperror.NotFoundException
		if errors.As(err, &nfe) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": nfe.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewDetailResponse(woman))
}
