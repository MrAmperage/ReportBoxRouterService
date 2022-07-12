package Models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"not null;username"`
	Password  string    `gorm:"not null;password"`
	Enabled   bool      `gorm:"not null;enabled"`
	RoleId    string    `gorm:"not null;rolename"`
	Role      Role      `gorm:"foreignkey:Id;references:RoleId"`
	StartDate time.Time `gorm:"not null;start_date"`
	EndDate   time.Time `gorm:"not null;end_date"`
}
