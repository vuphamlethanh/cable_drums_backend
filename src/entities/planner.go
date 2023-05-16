package entities

type Planner struct {
	AbstractBase
	AbstractAccount

	Contracts []Contract `gorm:"foreignKey:CreatedById"`
}

func NewPlanner(authId string) *Planner {
	return &Planner{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
