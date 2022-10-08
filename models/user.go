package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID           uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"omitempty,uuid"`
	Email            string    `json:"email" db:"email" redis:"email" validate:"required,email"`
	Password         string    `json:"password" db:"password" redis:"password" validate:"required,gte=8"`
	FirstName        string    `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName         string    `json:"last_name" db:"last_name" redis:"last_name" validate:"required,lte=30"`
	Role             string    `json:"role,omitempty" db:"role" redis:"role"`
	PhoneNumber      string    `json:"phone_number,omitempty" db:"phone_number" redis:"phone_number" validate:"omitempty,e164"`
	TotalParkingDays int       `json:"total_park_days" db:"total_park_days" redis:"total_park_days" validate:"gte=0"`
	CreatedAt        time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}
