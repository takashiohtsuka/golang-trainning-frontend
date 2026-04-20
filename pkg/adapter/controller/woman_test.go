package controller_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-trainning-frontend/pkg/apperror"
	"golang-trainning-frontend/pkg/collection"
	"golang-trainning-frontend/pkg/adapter/controller"
	responseWomen "golang-trainning-frontend/pkg/adapter/response/women"
	"golang-trainning-frontend/pkg/dto"
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
func (m *mockContext) Bind(i any) error          { return nil }
func (m *mockContext) Param(name string) string          { return m.params[name] }
func (m *mockContext) QueryParam(name string) string     { return m.queryParams[name] }
func (m *mockContext) Request() *http.Request            { return &http.Request{} }

// --- mock WomanUsecase ---

type mockWomanUsecase struct {
	getListInput  input.GetWomanListInput
	getListReturn collection.Collection[dto.WomanDTO]
	getListError  error

	getDetailInput  input.GetWomanDetailInput
	getDetailReturn dto.WomanDTO
	getDetailError  error
}

func (m *mockWomanUsecase) GetList(_ context.Context, i input.GetWomanListInput) (collection.Collection[dto.WomanDTO], error) {
	m.getListInput = i
	return m.getListReturn, m.getListError
}

func (m *mockWomanUsecase) GetDetail(i input.GetWomanDetailInput) (dto.WomanDTO, error) {
	m.getDetailInput = i
	return m.getDetailReturn, m.getDetailError
}

// --- GetWomanList tests ---

func TestWomanController_GetWomanList_WithNoStoreID_Returns200(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[dto.WomanDTO](nil),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	assert.Nil(t, uc.getListInput.StoreID)
}

func TestWomanController_GetWomanList_WithValidStoreID_PassesStoreIDToUsecase(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[dto.WomanDTO](nil),
	}
	ctx := &mockContext{queryParams: map[string]string{"store_id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	require.NotNil(t, uc.getListInput.StoreID)
	assert.Equal(t, uint(1), *uc.getListInput.StoreID)
}

func TestWomanController_GetWomanList_WithInvalidStoreID_Returns400(t *testing.T) {
	uc := &mockWomanUsecase{}
	ctx := &mockContext{queryParams: map[string]string{"store_id": "invalid"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, ctx.statusCode)
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
		getDetailReturn: &dto.Woman{ID: 1, Name: "女性1"},
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	assert.Equal(t, uint(1), uc.getDetailInput.WomanID)
}

func TestWomanController_GetWomanDetail_WithInvalidID_Returns400(t *testing.T) {
	uc := &mockWomanUsecase{}
	ctx := &mockContext{params: map[string]string{"id": "invalid"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, ctx.statusCode)
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
		getListReturn: collection.NewCollection[dto.WomanDTO]([]dto.WomanDTO{
			&dto.Woman{
				ID:     1,
				Name:   "女性1",
				Stores: collection.NewCollection[dto.WomanStore](nil),
				Images: collection.NewCollection[dto.WomanImage](nil),
				Blogs:  collection.NewCollection[dto.BlogDTO](nil),
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
		getDetailReturn: &dto.Woman{
			ID:     1,
			Name:   "女性1",
			Stores: collection.NewCollection[dto.WomanStore](nil),
			Images: collection.NewCollection[dto.WomanImage](nil),
			Blogs:  collection.NewCollection[dto.BlogDTO](nil),
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
