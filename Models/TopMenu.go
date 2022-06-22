package Models

type TopMenu struct {
	Caption  string `gorm:"primary_key;not null;caption"`
	Id       string `gorm:"not null;primary_key;id"`
	LeftMenu []LeftMenu
}
