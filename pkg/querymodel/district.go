package querymodel

type DistrictQueryModel interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetPrefectureID() uint
}

type District struct {
	ID           uint
	Name         string
	PrefectureID uint
}

func (d *District) IsNil() bool           { return d.ID == 0 }
func (d *District) GetID() uint           { return d.ID }
func (d *District) GetName() string       { return d.Name }
func (d *District) GetPrefectureID() uint { return d.PrefectureID }

type NilDistrict struct{}

func (n *NilDistrict) IsNil() bool           { return true }
func (n *NilDistrict) GetID() uint           { return 0 }
func (n *NilDistrict) GetName() string       { return "" }
func (n *NilDistrict) GetPrefectureID() uint { return 0 }
