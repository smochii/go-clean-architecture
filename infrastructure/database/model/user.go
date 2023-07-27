package model

import (
	"time"

	"github.com/smochii/go-clean-architecture/domain/value"
)

type User struct {
	Id        value.UserId         `gorm:"primaryKey;type:uuid"`
	Email     value.Email          `gorm:"not null;unique"`
	Password  value.HashedPassword `gorm:"not null"`
	CreatedAt time.Time            `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time            `gorm:"not null;default:current_timestamp"`
}
