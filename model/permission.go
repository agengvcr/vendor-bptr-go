package model

type Permission struct {
	Id   uint16 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
