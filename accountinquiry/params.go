package accountinquiry

import oygo "github.com/nadhirfr/oy-go"

// GetParams contains parameters for Get
type GetParams struct {
	BankCode      oygo.BankCode `form:"bank_code" json:"bank_code" validate:"required"`
	AccountNumber string        `form:"account_number" json:"account_number" validate:"required"`
}
