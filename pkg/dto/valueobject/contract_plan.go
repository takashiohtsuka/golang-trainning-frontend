package valueobject


type ContractPlan struct {
	ValueObject[string]
}

func NewContractPlan(code string) ContractPlan {
	return ContractPlan{NewValueObject(code)}
}

func EmptyContractPlan() ContractPlan {
	return ContractPlan{}
}

func (cp ContractPlan) GetCode() string {
	return cp.Get()
}

func (cp ContractPlan) IsEmpty() bool {
	return cp.Get() == ""
}
