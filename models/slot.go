package models

import (
	"time"

	"github.com/google/uuid"
)

type SlotSize string

const (
	Regular SlotSize = "regular"
	Large   SlotSize = "large"
)

type Slot struct {
	SlotID      uuid.UUID `json:"slot_id" db:"slot_id" redis:"slot_id" validate:"omitempty,uuid"`
	ParkingName string    `json:"parking_name" db:"parking_name" redis:"parking_name" validate:"required"`
	Number      int       `json:"number" db:"number" redis:"number" validate:"required, gt=0"`
	IsAvailable bool      `json:"is_available" db:"is_available" redis:"is_available" validate:"required"`
	PriceDay    string    `json:"price_day" db:"price_day" redis:"price_day" validate:"required,number"`
	Size        SlotSize  `json:"size" db:"size" redis:"size" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

type SlotSession struct {
	SlotSessionID uuid.UUID `json:"slot_session_id" db:"slot_session_id" redis:"slot_session_id" validate:"omitempty,uuid"`
	UserID        uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"require,uuid"`
	OrderID       uuid.UUID `json:"order_id" db:"order_id" redis:"order_id" validate:"omitempty,uuid"`
	CarNumber     string    `json:"car_number" db:"car_number" redis:"car_number" validate:"require"`
	StartedAt     time.Time `json:"started_at" db:"started_at" redis:"started_at" validate:"require"`
	FinishedAt    time.Time `json:"finished_at" db:"finished_at" redis:"finished_at" validate:"require"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
}
