package model

type Setting struct {
	Id    int    `gorm:"column:settingid"`
	Key   string `gorm:"column:settingkey"`
	Value string `gorm:"column:settingvalue"`
}
