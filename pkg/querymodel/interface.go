package querymodel

import (
	"golang-trainning-frontend/pkg/collection"
	fvo "golang-trainning-frontend/pkg/querymodel/valueobject"
)

type StoreQueryModel interface {
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
	GetWomen() collection.Collection[WomanQueryModel]
}

type WomanQueryModel interface {
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
	GetBlogs() collection.Collection[BlogQueryModel]
}

type BlogQueryModel interface {
	IsNil() bool
	GetID() uint
	GetWomanID() uint
	GetTitle() string
	GetBody() *string
	GetIsPublished() bool
	GetPhotos() collection.Collection[Photo]
}
