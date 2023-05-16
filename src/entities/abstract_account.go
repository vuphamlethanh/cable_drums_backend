package entities

type AbstractAccount struct {
	AuthId string `gorm:"unique;varchar"`
}

func NewAbstractAccount(authId string) AbstractAccount {
	return AbstractAccount{AuthId: authId}
}
