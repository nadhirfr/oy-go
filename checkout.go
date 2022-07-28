package oygo

import (
	"github.com/a-h/date"
)

type CheckoutResult struct {
	PaymentLinkId string `json:"payment_link_id"`
	Message       string `json:"message"`
	EmailStatus   string `json:"email_status"`
	Url           string `json:"url"`
	Status        bool   `json:"status"`
}

type CheckoutGet struct {
	Success    bool      `json:"success"`
	StatusCode int       `json:"status_code"`
	Reason     *string   `json:"reason"`
	Error      *string   `json:"error"`
	Data       *Checkout `json:"data"`
}

type Checkout struct {
	PartnerTxId       string               `json:"partner_tx_id"`
	Amount            int                  `json:"amount"`
	PaidAmount        *int                 `json:"paid_amount"`
	SenderName        *string              `json:"sender_name"`
	SenderPhoneNumber *string              `json:"sender_phone"`
	SenderNote        *string              `json:"sender_note"`
	SenderBank        *string              `json:"sender_bank"`
	SettlementStatus  *string              `json:"settlement_status"`
	SettlementType    *string              `json:"settlement_type"`
	SettlementTime    *date.YYYYMMDDHHMMSS `json:"settlement_time"`
	PaymentMethod     *string              `json:"payment_method"`
	Status            string               `json:"status"`
	IsInvoice         bool                 `json:"is_invoice"`
	TxRefNumber       *string              `json:"tx_ref_number"`
	Description       string               `json:"description"`
	Email             string               `json:"email"`
	Expiration        *date.YYYYMMDDHHMMSS `json:"expiration"`
	Created           date.YYYYMMDDHHMMSS  `json:"created"`
	Updated           date.YYYYMMDDHHMMSS  `json:"updated"`
}
