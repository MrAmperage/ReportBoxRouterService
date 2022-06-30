package Models

import "github.com/google/uuid"

type Organization struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption  string    `gorm:"not null;caption"`
	Internal bool      `gorm:"not null;internal"`
}
