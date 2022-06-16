package Models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Scheme struct {
	Id      uuid.UUID    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string       `gorm:"not null;caption"`
	Scheme  pgtype.JSONB `gorm:"not null;scheme"`
}
