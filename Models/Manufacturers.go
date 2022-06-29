package Models

import "github.com/google/uuid"

type Manufacturer struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
