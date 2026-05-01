package business_types

import "golang-trainning-frontend/pkg/querymodel"

type BusinessTypeItem struct {
	Code string `json:"code"`
}

type ListResponse struct {
	BusinessTypes []BusinessTypeItem `json:"business_types"`
}

func NewListResponse(businessTypes []querymodel.BusinessTypeQueryModel) ListResponse {
	items := make([]BusinessTypeItem, 0, len(businessTypes))
	for _, bt := range businessTypes {
		items = append(items, BusinessTypeItem{
			Code: bt.GetCode(),
		})
	}
	return ListResponse{BusinessTypes: items}
}
