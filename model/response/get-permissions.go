package response

import (
	"fmt"

	// "vendor.bptr.co.id/api-auth/lib"
	"vendor.bptr.co.id/api-auth/model"
)

type GetPermissionsResponse struct {
	EbEdit                bool `json:"expense_bill_edit"`
	EbDocumentReceive     bool `json:"expense_bill_documentreceive"`
	EbDocumentUnreceive   bool `json:"expense_bill_documentunreceive"`
	EbEtaxSave            bool `json:"expense_bill_etaxsave"`
	EbEtaxRevise          bool `json:"expense_bill_etaxrevise"`
	EtaxForceValid        bool `json:"etax_forcevalid"`
	UnitEdit              bool `json:"unit_edit"`
	UnitEtaxSave          bool `json:"unit_etaxsave"`
	UnitEtaxRevise        bool `json:"unit_etaxrevise"`
	PettyCashEdit         bool `json:"petty_cash_edit"`
	PettyCashEtaxSave     bool `json:"petty_cash_etaxsave"`
	PettycashEtaxRevise   bool `json:"petty_cash_etaxrevise"`
	CashAdvanceEdit       bool `json:"cash_advance_edit"`
	CashAdvanceEtaxSave   bool `json:"cash_advance_etaxsave"`
	CashAdvanceEtaxRevise bool `json:"cash_advance_etaxrevise"`
}

func NewGetPermissionsFromPermissions(permissions []model.Permission) GetPermissionsResponse {
	allPermission := make([]string, len(permissions))
	for i := 0; i < len(permissions); i++ {
		allPermission = append(allPermission, permissions[i].Name)
	}
	fmt.Println(allPermission)
	return GetPermissionsResponse{
		EbEdit:                lib.InArray("expenseBill_edit", allPermission),
		EbDocumentReceive:     lib.InArray("expenseBill_documentReceive", allPermission),
		EbDocumentUnreceive:   lib.InArray("expenseBill_documentUnreceive", allPermission),
		EbEtaxSave:            lib.InArray("expenseBill_etaxSave", allPermission),
		EbEtaxRevise:          lib.InArray("expenseBill_etaxRevise", allPermission),
		UnitEdit:              lib.InArray("unit_edit", allPermission),
		UnitEtaxSave:          lib.InArray("unit_etaxSave", allPermission),
		UnitEtaxRevise:        lib.InArray("unit_etaxRevise", allPermission),
		PettyCashEdit:         lib.InArray("pettyCash_edit", allPermission),
		PettyCashEtaxSave:     lib.InArray("pettyCash_etaxSave", allPermission),
		PettycashEtaxRevise:   lib.InArray("pettyCash_etaxRevise", allPermission),
		CashAdvanceEdit:       lib.InArray("cashAdvance_edit", allPermission),
		CashAdvanceEtaxSave:   lib.InArray("cashAdvance_etaxSave", allPermission),
		CashAdvanceEtaxRevise: lib.InArray("cashAdvance_etaxRevise", allPermission),
		EtaxForceValid:        lib.InArray("eTax_forceValid", allPermission),
	}
}
