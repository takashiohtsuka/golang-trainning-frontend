package prefectures

import "golang-trainning-frontend/pkg/querymodel"

type PrefectureItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ListResponse struct {
	Prefectures []PrefectureItem `json:"prefectures"`
}

func NewListResponse(prefectures []querymodel.PrefectureQueryModel) ListResponse {
	items := make([]PrefectureItem, 0, len(prefectures))
	for _, p := range prefectures {
		items = append(items, PrefectureItem{
			ID:   p.GetID(),
			Name: p.GetName(),
		})
	}
	return ListResponse{Prefectures: items}
}
