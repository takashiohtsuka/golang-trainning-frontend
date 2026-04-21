package querymodel

import (
	"time"

	"golang-trainning-frontend/pkg/collection"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
)

type OpenStatus string

const (
	OpenStatusOpen   OpenStatus = "open"
	OpenStatusClosed OpenStatus = "closed"
)

type Store struct {
	ID           uint                               `json:"id"`
	CompanyID    uint                               `json:"company_id"`
	District     fvo.District                       `json:"district"`
	Prefecture   fvo.Prefecture                     `json:"prefecture"`
	Region       fvo.Region                         `json:"region"`
	BusinessType fvo.BusinessType                   `json:"business_type"`
	ContractPlan fvo.ContractPlan                   `json:"contract_plan"`
	Name         string                             `json:"name"`
	IsActive     bool                               `json:"is_active"`
	OpenStatus   OpenStatus                         `json:"open_status"`
	Women        collection.Collection[WomanQueryModel] `json:"women"`
	CreatedAt    *time.Time                         `json:"created_at"`
	UpdatedAt    *time.Time                         `json:"updated_at"`
	DeletedAt    *time.Time                         `json:"deleted_at"`
}

func (s *Store) IsNil() bool                                    { return s.ID == 0 }
func (s *Store) GetID() uint                                     { return s.ID }
func (s *Store) GetCompanyID() uint                              { return s.CompanyID }
func (s *Store) GetDistrict() fvo.District                       { return s.District }
func (s *Store) GetPrefecture() fvo.Prefecture                   { return s.Prefecture }
func (s *Store) GetRegion() fvo.Region                           { return s.Region }
func (s *Store) GetBusinessType() fvo.BusinessType               { return s.BusinessType }
func (s *Store) GetContractPlan() fvo.ContractPlan               { return s.ContractPlan }
func (s *Store) GetName() string                                 { return s.Name }
func (s *Store) GetIsActive() bool                               { return s.IsActive }
func (s *Store) GetOpenStatus() OpenStatus                       { return s.OpenStatus }
func (s *Store) GetWomen() collection.Collection[WomanQueryModel] { return s.Women }

type NilStore struct{}

func (n *NilStore) IsNil() bool                                  { return true }
func (n *NilStore) GetID() uint                                  { return 0 }
func (n *NilStore) GetCompanyID() uint                           { return 0 }
func (n *NilStore) GetDistrict() fvo.District                    { return fvo.EmptyDistrict() }
func (n *NilStore) GetPrefecture() fvo.Prefecture                { return fvo.EmptyPrefecture() }
func (n *NilStore) GetRegion() fvo.Region                        { return fvo.EmptyRegion() }
func (n *NilStore) GetBusinessType() fvo.BusinessType            { return fvo.EmptyBusinessType() }
func (n *NilStore) GetContractPlan() fvo.ContractPlan            { return fvo.EmptyContractPlan() }
func (n *NilStore) GetName() string                              { return "" }
func (n *NilStore) GetIsActive() bool                            { return false }
func (n *NilStore) GetOpenStatus() OpenStatus                    { return "" }
func (n *NilStore) GetWomen() collection.Collection[WomanQueryModel] {
	return collection.NewCollection[WomanQueryModel](nil)
}
