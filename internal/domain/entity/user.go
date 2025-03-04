package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;uniqueIndex:user_unique"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null;uniqueIndex:user_unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
}
