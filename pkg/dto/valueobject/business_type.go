package valueobject


type BusinessType struct {
	ValueObject[string]
}

func NewBusinessType(code string) BusinessType {
	return BusinessType{NewValueObject(code)}
}

func EmptyBusinessType() BusinessType {
	return BusinessType{}
}

func (bt BusinessType) GetCode() string {
	return bt.Get()
}

func (bt BusinessType) IsEmpty() bool {
	return bt.Get() == ""
}
