package repository

import (
	"database/sql"

	"gorm.io/gorm"
	"vendor.bptr.co.id/api-auth/model"
)

type VendorUserRepository struct {
	db *gorm.DB
}

func NewVendorUserRepository(db *gorm.DB) *VendorUserRepository {
	return &VendorUserRepository{
		db: db,
	}
}

func (ur *VendorUserRepository) GetUserById(id int) model.VendorUser {
	var vendorUser model.VendorUser
	res := ur.db.Raw(""+
		"SELECT vendoruserid, vendorusername, vendoruserpassword, vendoruserlastlogin, vendoruserfullname FROM \"vendor\" "+
		"WHERE vendoruseractive = '1'"+
		"AND vendoruserid = @vendoruserid ",
		sql.Named("vendoruserid", id)).Take(&vendorUser)
	if res.Error != nil {
		return vendorUser
	}

	return vendorUser
}

func (ur *VendorUserRepository) GetUserByUsername(username string) model.VendorUser {
	var vendorUser model.VendorUser
	res := ur.db.Raw(""+
		"SELECT vendoruserid, vendorusername, vendoruserpassword, vendoruserlastlogin, vendoruserfullname FROM \"vendor\" "+
		"WHERE vendoruseractive = '1'"+
		"AND vendorusername = @username ",
		sql.Named("username", username)).Take(&vendorUser)
	if res.Error != nil {
		return vendorUser
	}

	return vendorUser
}
