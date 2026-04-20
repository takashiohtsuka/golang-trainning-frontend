package dto

import (
	"golang-trainning-frontend/pkg/collection"
	fvo "golang-trainning-frontend/pkg/dto/valueobject"
)

type StoreDTO interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetDistrict() fvo.District
	GetPrefecture() fvo.Prefecture
	GetRegion() fvo.Region
	GetBusinessType() fvo.BusinessType
	GetContractPlan() fvo.ContractPlan
	GetName() string
	GetIsActive() bool
	GetOpenStatus() OpenStatus
	GetWomen() collection.Collection[WomanDTO]
}

type WomanDTO interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetName() string
	GetAge() *int
	GetBirthplace() *string
	GetBloodType() *string
	GetHobby() *string
	GetIsActive() bool
	GetStores() collection.Collection[WomanStore]
	GetImages() collection.Collection[WomanImage]
	GetBlogs() collection.Collection[BlogDTO]
}

type BlogDTO interface {
	IsNil() bool
	GetID() uint
	GetWomanID() uint
	GetTitle() string
	GetBody() *string
	GetIsPublished() bool
	GetPhotos() collection.Collection[Photo]
}
