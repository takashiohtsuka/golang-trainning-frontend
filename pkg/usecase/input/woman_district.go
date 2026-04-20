package input

type GetWomanDistrictListInput struct {
	DistrictID uint
	Page       uint
	BloodTypes []string
	AgeRanges  []string
}
