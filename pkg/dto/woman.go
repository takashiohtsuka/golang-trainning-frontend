package dto

import (
	"time"

	"golang-trainning-frontend/pkg/collection"
)

type Woman struct {
	ID               uint                                        `json:"id"`
	CompanyID        uint                                        `json:"company_id"`
	Name             string                                      `json:"name"`
	Age              *int                                        `json:"age"`
	Birthplace       *string                                     `json:"birthplace"`
	BloodType        *string                                     `json:"blood_type"`
	Hobby            *string                                     `json:"hobby"`
	IsActive         bool                                        `json:"is_active"`
	Stores           collection.Collection[WomanStore]           `json:"stores"`
	Images           collection.Collection[WomanImage]           `json:"images"`
	Blogs            collection.Collection[BlogDTO]           `json:"blogs"`
	CreatedAt        *time.Time                                  `json:"created_at"`
	UpdatedAt        *time.Time                                  `json:"updated_at"`
	DeletedAt        *time.Time                                  `json:"deleted_at"`
}

func (w *Woman) IsNil() bool               { return w.ID == 0 }
func (w *Woman) GetID() uint               { return w.ID }
func (w *Woman) GetCompanyID() uint        { return w.CompanyID }
func (w *Woman) GetName() string           { return w.Name }
func (w *Woman) GetAge() *int              { return w.Age }
func (w *Woman) GetBirthplace() *string    { return w.Birthplace }
func (w *Woman) GetBloodType() *string     { return w.BloodType }
func (w *Woman) GetHobby() *string         { return w.Hobby }
func (w *Woman) GetIsActive() bool         { return w.IsActive }
func (w *Woman) GetStores() collection.Collection[WomanStore] { return w.Stores }
func (w *Woman) GetImages() collection.Collection[WomanImage] { return w.Images }
func (w *Woman) GetBlogs() collection.Collection[BlogDTO]     { return w.Blogs }

type NilWoman struct{}

func (n *NilWoman) IsNil() bool                               { return true }
func (n *NilWoman) GetID() uint                               { return 0 }
func (n *NilWoman) GetCompanyID() uint                        { return 0 }
func (n *NilWoman) GetName() string                           { return "" }
func (n *NilWoman) GetAge() *int                              { return nil }
func (n *NilWoman) GetBirthplace() *string                    { return nil }
func (n *NilWoman) GetBloodType() *string                     { return nil }
func (n *NilWoman) GetHobby() *string                         { return nil }
func (n *NilWoman) GetIsActive() bool                         { return false }
func (n *NilWoman) GetStores() collection.Collection[WomanStore] {
	return collection.NewCollection[WomanStore](nil)
}
func (n *NilWoman) GetImages() collection.Collection[WomanImage] {
	return collection.NewCollection[WomanImage](nil)
}
func (n *NilWoman) GetBlogs() collection.Collection[BlogDTO] {
	return collection.NewCollection[BlogDTO](nil)
}
