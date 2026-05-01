package women

import (
	"time"

	"golang-trainning-frontend/pkg/querymodel"
)

type ImmediateAvailableListResponse struct {
	Women []ImmediateAvailableWomanItem `json:"women"`
	Total uint                          `json:"total"`
}

type ImmediateAvailableWomanItem struct {
	ID           uint        `json:"id"`
	Name         string      `json:"name"`
	Age          *int        `json:"age"`
	Birthplace   *string     `json:"birthplace"`
	BloodType    *string     `json:"blood_type"`
	Hobby        *string     `json:"hobby"`
	Store        StoreItem   `json:"store"`
	Images       []ImageItem `json:"images"`
	ExpiresAt    *time.Time  `json:"expires_at"`
}

func NewImmediateAvailableListResponse(women []querymodel.ImmediateAvailableWomanQueryModel, total uint) ImmediateAvailableListResponse {
	items := make([]ImmediateAvailableWomanItem, 0, len(women))
	for _, w := range women {
		items = append(items, toImmediateAvailableWomanItem(w))
	}
	return ImmediateAvailableListResponse{Women: items, Total: total}
}

func toImmediateAvailableWomanItem(w querymodel.ImmediateAvailableWomanQueryModel) ImmediateAvailableWomanItem {
	store := w.GetStore()

	images := make([]ImageItem, 0)
	for _, i := range w.GetImages().All() {
		images = append(images, ImageItem{
			ID:   i.ID,
			Path: i.Path,
		})
	}

	return ImmediateAvailableWomanItem{
		ID:         w.GetID(),
		Name:       w.GetName(),
		Age:        w.GetAge(),
		Birthplace: w.GetBirthplace(),
		BloodType:  w.GetBloodType(),
		Hobby:      w.GetHobby(),
		Store: StoreItem{
			ID:           store.ID,
			Name:         store.Name,
			BusinessType: store.BusinessType.GetCode(),
		},
		Images:    images,
		ExpiresAt: w.GetExpiresAt(),
	}
}
