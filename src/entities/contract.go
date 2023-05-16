package entities

import (
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	AbstractBase

	StartDay      time.Time
	EndDay        time.Time
	CableQuantity uint

	SupplierId  uuid.UUID
	CreatedById uuid.UUID

	Supplier  *Supplier
	CreatedBy *Planner
}
