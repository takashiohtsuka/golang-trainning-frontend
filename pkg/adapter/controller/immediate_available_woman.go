package controller

import (
	"net/http"

	requestIAW "golang-trainning-frontend/pkg/adapter/request/immediate_available_women"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/usecase/inputport"
)

type immediateAvailableWomanController struct {
	usecase inputport.ImmediateAvailableWomanUsecase
}

type ImmediateAvailableWoman interface {
	GetImmediateAvailableWomanList(c Context) error
}

func NewImmediateAvailableWomanController(u inputport.ImmediateAvailableWomanUsecase) ImmediateAvailableWoman {
	return &immediateAvailableWomanController{usecase: u}
}

func (ic *immediateAvailableWomanController) GetImmediateAvailableWomanList(ctx Context) error {
	var req requestIAW.ListRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	women, total, err := ic.usecase.GetList(ctx.Request().Context(), req.ToInput())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, responseWomen.NewImmediateAvailableListResponse(women.All(), total))
}
