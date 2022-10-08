package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentProvider struct {
	PaymentProviderID uuid.UUID `json:"payment_provider_id" db:"payment_provider_id" redis:"payment_provider_id" validate:"omitempty,uuid"`
	UserID            uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"required,uuid"`
	AccountNumber     string    `json:"account_number" db:"account_number" redis:"account_number" validate:"required,number,eq=16"`
	ExpiryDate        time.Time `json:"exp_date" db:"exp_date" redis:"exp_date" validate:"required"`
	ProviderName      string    `json:"provider_name,omitempty" db:"provider_name" redis:"provider_name" validate:"omitempty,lte=30"`
	CreatedAt         time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}
