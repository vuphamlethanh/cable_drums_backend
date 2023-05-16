package entities

type Contractor struct {
	AbstractBase
	AbstractAccount
}

func NewContractor(authId string) *Contractor {
	return &Contractor{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
