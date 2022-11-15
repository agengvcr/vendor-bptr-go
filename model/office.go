package model

type Office struct {
	Id         int    `gorm:"column:officeid"`
	Name       string `gorm:"column:officename"`
	Address    string `gorm:"column:officeaddress"`
	City       string `gorm:"column:officecity"`
	Province   string `gorm:"column:officeprovince"`
	Email      string `gorm:"column:officeemail"`
	CpName     string `gorm:"column:officecpname"`
	CpPhone    string `gorm:"column:officecpphone"`
	TaxNumber  string `gorm:"column:officetaxnumber"`
	TaxName    string `gorm:"column:officetaxname"`
	TaxAddress string `gorm:"column:officetaxaddress"`
}
