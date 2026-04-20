package women

import "golang-trainning-frontend/pkg/dto"

type DistrictListResponse struct {
	Women []WomanListItem `json:"women"`
	Total uint            `json:"total"`
}

func NewDistrictListResponse(women []dto.WomanDTO, total uint) DistrictListResponse {
	items := make([]WomanListItem, 0, len(women))
	for _, w := range women {
		items = append(items, toWomanListItem(w))
	}
	return DistrictListResponse{Women: items, Total: total}
}
