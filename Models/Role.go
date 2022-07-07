package Models

import "github.com/google/uuid"

type Role struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Rolename string    `gorm:"not null;rolename"`
}
