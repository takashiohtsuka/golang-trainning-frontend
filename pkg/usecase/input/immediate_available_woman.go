package input

type GetImmediateAvailableWomanListInput struct {
	Page          uint
	Limit         uint
	PrefectureID  uint
	DistrictID    uint
	BusinessTypes []string
	BloodTypes    []string
	AgeRanges     []string
}
