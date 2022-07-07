package Models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username string    `gorm:"primaryKey;not null;username"`
	Password string    `gorm:"not null;password"`
	Enabled  bool      `gorm:"not null;enabled"`
	RoleId   string    `gorm:"not null;rolename"`
	Role     Role      `gorm:"foreignkey:Id;references:role_id"`
}
