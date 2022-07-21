package Models

import "github.com/google/uuid"

type Area struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4();id"`
	Caption string    `gorm:"not null;caption"`
}
