package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email      string    `gorm:"unique;not null"`
	TOTPSecret string    `json:"-"`
	IsVerified bool
}
