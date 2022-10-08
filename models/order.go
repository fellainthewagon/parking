package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID    uuid.UUID `json:"order_id" db:"order_id" redis:"order_id" validate:"omitempty,uuid"`
	UserID     uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"require,uuid"`
	SlotID     uuid.UUID `json:"slot_id" db:"slot_id" redis:"slot_id" validate:"require,uuid"`
	Days       int       `json:"days" db:"days" redis:"days" validate:"required"`
	Total      string    `json:"total" db:"total" redis:"total" validate:"required,number,gt=0"`
	Discount   string    `json:"discount" db:"discount" redis:"discount" validate:"required,number,gte=0"`
	GrandTotal string    `json:"grand_total" db:"grand_total" redis:"grand_total" validate:"required,number,gt=0"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}
