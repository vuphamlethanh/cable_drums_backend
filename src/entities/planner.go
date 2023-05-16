package entities

type Planner struct {
	AbstractBase
	AbstractAccount
}

func NewPlanner(authId string) *Planner {
	return &Planner{AbstractBase: NewAbstractBase(), AbstractAccount: NewAbstractAccount(authId)}
}
