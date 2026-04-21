package controller_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-trainning-frontend/pkg/adapter/controller"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/apperror"
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/querymodel"
	"golang-trainning-frontend/pkg/usecase/input"
	"golang-trainning-frontend/pkg/usecase/query"
)

// --- mock Context ---

type mockContext struct {
	params      map[string]string
	queryParams map[string]string
	statusCode  int
	body        any
}

func (m *mockContext) JSON(code int, i any) error {
	m.statusCode = code
	m.body = i
	return nil
}
func (m *mockContext) Bind(i any) error       { return nil }
func (m *mockContext) Validate(i any) error   { return nil }
func (m *mockContext) Param(name string) string      { return m.params[name] }
func (m *mockContext) QueryParam(name string) string { return m.queryParams[name] }
func (m *mockContext) Request() *http.Request        { return &http.Request{} }

// --- mock WomanUsecase ---

type mockWomanUsecase struct {
	getListReturn collection.Collection[querymodel.WomanQueryModel]
	getListError  error

	getStoreWomanListInput  input.GetStoreWomanListInput
	getStoreWomanListReturn collection.Collection[querymodel.WomanQueryModel]
	getStoreWomanListError  error

	getDetailInput  input.GetWomanDetailInput
	getDetailReturn querymodel.WomanQueryModel
	getDetailError  error
}

func (m *mockWomanUsecase) GetList(_ context.Context, _ input.GetWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error) {
	return m.getListReturn, m.getListError
}

func (m *mockWomanUsecase) GetStoreWomanList(_ context.Context, i input.GetStoreWomanListInput) (collection.Collection[querymodel.WomanQueryModel], error) {
	m.getStoreWomanListInput = i
	return m.getStoreWomanListReturn, m.getStoreWomanListError
}

func (m *mockWomanUsecase) GetDetail(_ context.Context, i input.GetWomanDetailInput) (querymodel.WomanQueryModel, error) {
	m.getDetailInput = i
	return m.getDetailReturn, m.getDetailError
}

// --- GetWomanList tests ---

func TestWomanController_GetWomanList_Returns200(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[querymodel.WomanQueryModel](nil),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
}

func TestWomanController_GetWomanList_WhenUsecaseFails_Returns500(t *testing.T) {
	uc := &mockWomanUsecase{
		getListError: errors.New("db error"),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.statusCode)
}

// --- GetWomanDetail tests ---

func TestWomanController_GetWomanDetail_WithValidID_Returns200(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailReturn: &querymodel.Woman{ID: 1, Name: "女性1"},
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
}

func TestWomanController_GetWomanDetail_WhenNotFound_Returns404(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailError: apperror.NewNotFoundException("woman not found"),
	}
	ctx := &mockContext{params: map[string]string{"id": "99999"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, ctx.statusCode)
}

func TestWomanController_GetWomanDetail_WhenUsecaseFails_Returns500(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailError: errors.New("db error"),
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.statusCode)
}

// --- GetWomanList response body tests ---

func TestWomanController_GetWomanList_ResponseBodyIsListResponse(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[querymodel.WomanQueryModel]([]querymodel.WomanQueryModel{
			&querymodel.Woman{
				ID:     1,
				Name:   "女性1",
				Stores: collection.NewCollection[querymodel.WomanStore](nil),
				Images: collection.NewCollection[querymodel.WomanImage](nil),
				Blogs:  collection.NewCollection[querymodel.BlogQueryModel](nil),
			},
		}),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	body, ok := ctx.body.(responseWomen.ListResponse)
	require.True(t, ok, "body は ListResponse 型であること")
	require.Len(t, body.Women, 1)
	assert.Equal(t, uint(1), body.Women[0].ID)
	assert.Equal(t, "女性1", body.Women[0].Name)
}

// --- GetWomanDetail response body tests ---

func TestWomanController_GetWomanDetail_ResponseBodyIsDetailResponse(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailReturn: &querymodel.Woman{
			ID:     1,
			Name:   "女性1",
			Stores: collection.NewCollection[querymodel.WomanStore](nil),
			Images: collection.NewCollection[querymodel.WomanImage](nil),
			Blogs:  collection.NewCollection[querymodel.BlogQueryModel](nil),
		},
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	body, ok := ctx.body.(responseWomen.DetailResponse)
	require.True(t, ok, "body は DetailResponse 型であること")
	assert.Equal(t, uint(1), body.ID)
	assert.Equal(t, "女性1", body.Name)
}

// --- query.Condition の検証ヘルパー（参考） ---

var _ = query.Where // query パッケージが使われていることの確認
