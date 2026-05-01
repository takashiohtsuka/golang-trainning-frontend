package querymodel

type PrefectureQueryModel interface {
	IsNil() bool
	GetID() uint
	GetName() string
}

type Prefecture struct {
	ID   uint
	Name string
}

func (p *Prefecture) IsNil() bool    { return p.ID == 0 }
func (p *Prefecture) GetID() uint    { return p.ID }
func (p *Prefecture) GetName() string { return p.Name }

type NilPrefecture struct{}

func (n *NilPrefecture) IsNil() bool    { return true }
func (n *NilPrefecture) GetID() uint    { return 0 }
func (n *NilPrefecture) GetName() string { return "" }
