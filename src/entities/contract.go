package entities

import (
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	AbstractBase

	StartDay      time.Time
	EndDay        time.Time
	CableQuantity uint `gorm:"type:integer"`

	SupplierId  uuid.UUID
	CreatedById uuid.UUID

	Supplier  *Supplier
	CreatedBy *Planner

	WithdrawRequests []WithdrawRequest `gorm:"foreignKey:ContractId"`
}
