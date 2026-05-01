package immediate_available_women

import (
	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/usecase/input"
)

type ListRequest struct {
	Page         uint   `query:"page"          validate:"omitempty,min=1"`
	Limit        uint   `query:"limit"         validate:"omitempty,min=1"`
	PrefectureID uint   `query:"prefecture_id"`
	DistrictID   uint   `query:"district_id"`
	BusinessType string `query:"business_type"`
	BloodType    string `query:"blood_type"`
	AgeRange     string `query:"age_range"`
}

func (req *ListRequest) ToInput() input.GetImmediateAvailableWomanListInput {
	return input.GetImmediateAvailableWomanListInput{
		Page:          req.Page,
		Limit:         req.Limit,
		PrefectureID:  req.PrefectureID,
		DistrictID:    req.DistrictID,
		BusinessTypes: helper.SplitComma(req.BusinessType),
		BloodTypes:    helper.SplitComma(req.BloodType),
		AgeRanges:     helper.SplitComma(req.AgeRange),
	}
}
