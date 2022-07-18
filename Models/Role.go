package Models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Role struct {
	Id                  uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Rolename            string         `gorm:"not null;rolename"`
	MenusAccess         pq.StringArray `gorm:"type:varchar[];default:array[]::varchar[];menus_access"`
	OrganizationsAccess pq.StringArray `gorm:"type:varchar[];default:array[]::varchar[];organizations_access"`
}
