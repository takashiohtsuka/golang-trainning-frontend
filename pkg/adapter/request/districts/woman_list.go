package districts

import (
	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanListRequest struct {
	ID        uint   `param:"id"         validate:"required"`
	Page      uint   `query:"page"       validate:"omitempty,min=1"`
	BloodType string `query:"blood_type"`
	AgeRange  string `query:"age_range"`
}

func (req *WomanListRequest) ToInput() input.GetWomanDistrictListInput {
	return input.GetWomanDistrictListInput{
		DistrictID: req.ID,
		Page:       req.Page,
		BloodTypes: helper.SplitComma(req.BloodType),
		AgeRanges:  helper.SplitComma(req.AgeRange),
	}
}
