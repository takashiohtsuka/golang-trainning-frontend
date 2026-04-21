package districts

import (
	"golang-trainning-frontend/pkg/helper"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanCountRequest struct {
	ID        uint   `param:"id"         validate:"required"`
	BloodType string `query:"blood_type"`
	AgeRange  string `query:"age_range"`
}

func (req *WomanCountRequest) ToInput() input.GetWomanDistrictCountInput {
	return input.GetWomanDistrictCountInput{
		DistrictID: req.ID,
		BloodTypes: helper.SplitComma(req.BloodType),
		AgeRanges:  helper.SplitComma(req.AgeRange),
	}
}
