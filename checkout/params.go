package checkout

import "time"

type CreateParams struct {
	Description               string     `json:"description"`
	PartnerTxId               string     `json:"partner_tx_id"`
	Notes                     string     `json:"notes"`
	SenderName                string     `json:"sender_name"`
	Amount                    int        `json:"amount"`
	Email                     string     `json:"email"`
	PhoneNumber               string     `json:"phone_number"`
	IsOpen                    bool       `json:"is_open"`
	Step                      *string    `json:"step"`
	IncludeAdminFee           bool       `json:"include_admin_fee"`
	LisDisabledPaymentMethods string     `json:"list_disabled_payment_methods"`
	ListEnabledBanks          string     `json:"list_enabled_banks"`
	ListEnabledEwallet        string     `json:"list_enabled_ewallet"`
	Expiration                *time.Time `json:"expiration"`
	DueDate                   *time.Time `json:"due_date"`
}

// GetParams contains parameters for Get
type GetParams struct {
	PartnerTxId string `json:"partner_tx_id" validate:"required"`
}
