package model

import (
	"time"
)

// User model structure
type VendorUser struct {
	Id        int       `gorm:"column:vendoruserid"`
	Username  string    `gorm:"column:vendorusername"`
	Password  string    `gorm:"column:vendoruserpassword"`
	LastLogin time.Time `gorm:"column:vendoruserlastlogin"`
	Name      string    `gorm:"column:vendoruserfullname"`
}
