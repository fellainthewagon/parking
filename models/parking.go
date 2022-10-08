package models

import (
	"time"

	"github.com/google/uuid"
)

type ParkingType string

const (
	Indoor  ParkingType = "indoor"
	Outdoor ParkingType = "outdoor"
)

type DiscountLevel string

const (
	Basic    DiscountLevel = "basic"
	Gold     DiscountLevel = "gold"
	Platinum DiscountLevel = "platinum"
)

type Parking struct {
	ParkingName string      `json:"parking_name" db:"parking_name" redis:"parking_name" validate:"required"`
	Type        ParkingType `json:"parking_type" db:"parking_type" redis:"parking_type" validate:"required"`
	Capacity    int         `json:"capacity" db:"capacity" redis:"capacity" validate:"required,gte=100"`
	City        string      `json:"city" db:"city" redis:"city" validate:"required"`
	AddressLine string      `json:"address_line" db:"address_line" redis:"address_line" validate:"required"`
	PhoneNumber string      `json:"phone_number" db:"phone_number" redis:"phone_number" validate:"required,e164"`
	Email       string      `json:"email" db:"email" redis:"email" validate:"required,email"`
	Description string      `json:"description" db:"description" redis:"description" validate:"required,lte=100"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

type DiscountPolicy struct {
	DiscPolicyID  uuid.UUID     `json:"disc_policy_id" db:"disc_policy_id" redis:"disc_policy_id" validate:"omitempty,uuid"`
	ParkingName   string        `json:"parking_name" db:"parking_name" redis:"parking_name" validate:"required"`
	Level         DiscountLevel `json:"level" db:"level" redis:"level" validate:"required"`
	ConditionDays int           `json:"condition_days" db:"condition_days" redis:"condition_days" validate:"required,gt=0"`
	Value         string        `json:"value" db:"value" redis:"value" validate:"required,number"`
	CreatedAt     time.Time     `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}
