package entities

type Contractor struct {
	AbstractBase
	AbstractAccount

	WithdrawRequests []WithdrawRequest `gorm:"foreignKey:ContractorId"`
}

func NewContractor(authId string) *Contractor {
	return &Contractor{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
