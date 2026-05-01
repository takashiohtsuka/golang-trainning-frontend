package querymodel

type BusinessTypeQueryModel interface {
	IsNil() bool
	GetCode() string
}

type BusinessType struct {
	Code string
}

func (bt *BusinessType) IsNil() bool     { return bt.Code == "" }
func (bt *BusinessType) GetCode() string { return bt.Code }

type NilBusinessType struct{}

func (n *NilBusinessType) IsNil() bool     { return true }
func (n *NilBusinessType) GetCode() string { return "" }
