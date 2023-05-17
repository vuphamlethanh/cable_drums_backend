package entities

import "github.com/google/uuid"

type WithdrawRequest struct {
	AbstractBase

	CableQuantity uint   `gorm:"type:integer"`
	Status        string `gorm:"type:varchar"`

	ContractId   uuid.UUID
	CreatedById  uuid.UUID
	ContractorId uuid.UUID

	Contract   *Contract
	CreatedBy  *Planner
	Contractor *Contractor
}
