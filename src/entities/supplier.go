package entities

type Supplier struct {
	AbstractBase
	AbstractAccount

	Contracts []Contract `gorm:"foreignKey:SupplierId"`
}

func NewSupplier(authId string) *Supplier {
	return &Supplier{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
