package model

type ApiKey struct {
	Secret string `gorm:"column:apikeysecret"`
	Token  string `gorm:"column:apikeytoken"`
	Type   string `gorm:"column:apikeytype"`
}
