package response

import (
	"vendor.bptr.co.id/api-auth/model"
)

type GetOfficeProfileResponse struct {
	OfficeId         int    `json:"id"`
	OfficeName       string `json:"name"`
	OfficeAddress    string `json:"address"`
	OfficeCity       string `json:"city"`
	OfficeProvince   string `json:"province"`
	OfficeEmail      string `json:"email"`
	OfficeCpName     string `json:"cp_name"`
	OfficeCpPhone    string `json:"cp_phone"`
	OfficeTaxNumber  string `json:"tax_number"`
	OfficeTaxName    string `json:"tax_name"`
	OfficeTaxAddress string `json:"tax_address"`
}

func NewGetOfficeTaxResponseFromOffice(office model.Office) GetOfficeProfileResponse {

	return GetOfficeProfileResponse{
		OfficeId:         office.Id,
		OfficeName:       office.Name,
		OfficeAddress:    office.Address,
		OfficeCity:       office.City,
		OfficeProvince:   office.Province,
		OfficeEmail:      office.Email,
		OfficeCpName:     office.CpName,
		OfficeCpPhone:    office.CpPhone,
		OfficeTaxNumber:  office.TaxNumber,
		OfficeTaxName:    office.TaxName,
		OfficeTaxAddress: office.TaxAddress,
	}
}
