package Models

import "github.com/jackc/pgtype"

type Role struct {
	Rolename    string       `gorm:"primaryKey;not null;rolename"`
	Permissions pgtype.JSONB `gorm:"not null;permissions"`
	Comment     string       `gorm:"not null;comment"`
}
