package Models

import "github.com/google/uuid"

type LeftMenu struct {
	Caption   string    `gorm:"primary_key;not null;caption"`
	Id        string    `gorm:"not null;id"`
	TopMenuId string    `gorm:"not null;top_menu_id"`
	SchemeId  uuid.UUID `gorm:"type:uuid;scheme_id"`
}
