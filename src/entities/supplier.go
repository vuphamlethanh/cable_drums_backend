package entities

type Supplier struct {
	AbstractBase
	AbstractAccount
}

func NewSupplier(authId string) *Supplier {
	return &Supplier{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
