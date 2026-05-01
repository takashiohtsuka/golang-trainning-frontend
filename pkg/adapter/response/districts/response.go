package districts

import "golang-trainning-frontend/pkg/querymodel"

type DistrictItem struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	PrefectureID uint   `json:"prefecture_id"`
}

type ListResponse struct {
	Districts []DistrictItem `json:"districts"`
}

func NewListResponse(districts []querymodel.DistrictQueryModel) ListResponse {
	items := make([]DistrictItem, 0, len(districts))
	for _, d := range districts {
		items = append(items, DistrictItem{
			ID:           d.GetID(),
			Name:         d.GetName(),
			PrefectureID: d.GetPrefectureID(),
		})
	}
	return ListResponse{Districts: items}
}
