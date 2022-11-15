package model

import (
	"time"
)

//User model structure
type User struct {
	Id        int       `gorm:"column:userid"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:userpassword"`
	LastLogin time.Time `gorm:"column:userlastlogin"`
	Email     string    `gorm:"column:useremail"`
	Name      string    `gorm:"column:userfullname"`
}
