package entities

import (
	"github.com/google/uuid"
	"time"
)

type AbstractBase struct {
	Id         uuid.UUID `gorm:"primaryKey;varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
}

func NewAbstractBase() AbstractBase {
	return AbstractBase{Id: uuid.New()}
}
