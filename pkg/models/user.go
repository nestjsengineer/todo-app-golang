package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"-"`
}
