package entities

type Admin struct {
	AbstractBase
	AbstractAccount
}

func NewAdmin(authId string) *Admin {
	return &Admin{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
