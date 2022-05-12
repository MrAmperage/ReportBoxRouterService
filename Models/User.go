package Models

type User struct {
	Username string `gorm:"primaryKey;not null;username"`
	Password string `gorm:"not null;password"`
	Enabled  bool   `gorm:"not null;enabled"`
	Rolename string `gorm:"not null;rolename"`
}
