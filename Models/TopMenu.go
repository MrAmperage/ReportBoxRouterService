package Models

import "github.com/google/uuid"

type TopMenu struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"caption"`
	Type    string
	Items   []SubMenu `gorm:"foreignkey:TopMenuId"`
}
