package controller

import (
	"net/http"

	requestStores "golang-trainning-frontend/pkg/adapter/request/stores"
	responseStores "golang-trainning-frontend/pkg/adapter/response/stores"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type storeController struct {
	storeUsecase inputport.StoreUsecase
}

type Store interface {
	GetStoreDetail(c Context) error
}

func NewStoreController(u inputport.StoreUsecase) Store {
	return &storeController{u}
}

func (sc *storeController) GetStoreDetail(ctx Context) error {
	var req requestStores.GetRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	store, err := sc.storeUsecase.GetDetail(ctx.Request().Context(), req.ToStoreDetailInput())
	if err != nil {
		return err
	}

	if store.IsNil() {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	return ctx.JSON(http.StatusOK, responseStores.NewDetailResponse(store))
}
