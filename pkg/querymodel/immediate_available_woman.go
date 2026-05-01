package querymodel

import (
	"time"

	"golang-trainning-frontend/pkg/collection"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
)

type ImmediateAvailableWomanQueryModel interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetAge() *int
	GetBirthplace() *string
	GetBloodType() *string
	GetHobby() *string
	GetStore() ImmediateAvailableWomanStore
	GetImages() collection.Collection[WomanImage]
	GetExpiresAt() *time.Time
}

type ImmediateAvailableWoman struct {
	ID         uint
	Name       string
	Age        *int
	Birthplace *string
	BloodType  *string
	Hobby      *string
	Store      ImmediateAvailableWomanStore
	Images     collection.Collection[WomanImage]
	ExpiresAt  *time.Time
}

func (w *ImmediateAvailableWoman) IsNil() bool                                          { return w.ID == 0 }
func (w *ImmediateAvailableWoman) GetID() uint                                          { return w.ID }
func (w *ImmediateAvailableWoman) GetName() string                                      { return w.Name }
func (w *ImmediateAvailableWoman) GetAge() *int                                         { return w.Age }
func (w *ImmediateAvailableWoman) GetBirthplace() *string                               { return w.Birthplace }
func (w *ImmediateAvailableWoman) GetBloodType() *string                                { return w.BloodType }
func (w *ImmediateAvailableWoman) GetHobby() *string                                    { return w.Hobby }
func (w *ImmediateAvailableWoman) GetStore() ImmediateAvailableWomanStore               { return w.Store }
func (w *ImmediateAvailableWoman) GetImages() collection.Collection[WomanImage]         { return w.Images }
func (w *ImmediateAvailableWoman) GetExpiresAt() *time.Time                             { return w.ExpiresAt }

type NilImmediateAvailableWoman struct{}

func (n *NilImmediateAvailableWoman) IsNil() bool                                          { return true }
func (n *NilImmediateAvailableWoman) GetID() uint                                          { return 0 }
func (n *NilImmediateAvailableWoman) GetName() string                                      { return "" }
func (n *NilImmediateAvailableWoman) GetAge() *int                                         { return nil }
func (n *NilImmediateAvailableWoman) GetBirthplace() *string                               { return nil }
func (n *NilImmediateAvailableWoman) GetBloodType() *string                                { return nil }
func (n *NilImmediateAvailableWoman) GetHobby() *string                                    { return nil }
func (n *NilImmediateAvailableWoman) GetStore() ImmediateAvailableWomanStore               { return ImmediateAvailableWomanStore{} }
func (n *NilImmediateAvailableWoman) GetImages() collection.Collection[WomanImage]         { return collection.NewCollection[WomanImage](nil) }
func (n *NilImmediateAvailableWoman) GetExpiresAt() *time.Time                             { return nil }

type ImmediateAvailableWomanStore struct {
	ID           uint             `json:"id"`
	Name         string           `json:"name"`
	BusinessType fvo.BusinessType `json:"business_type"`
}
