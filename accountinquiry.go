package oygo

import (
	"github.com/a-h/date"
)


type AccountInquiryGet struct {
	Status	AccountInquiryStatus	`json:"status"` //Status of inquiry {code: <status_code>, message: <status_message>}
	BankCode	BankCode	`json:"bank_code"` //Bank Code of the Beneficiary account, see Disbursement Bank Codes
	AccountNumber	string	`json:"account_number"` //Account Number of the Beneficiary Account
	AccountName	string	`json:"account_name"` //Account Name of the Beneficiary Account
	Id	string	`json:"id"` //Unique ID of the inquiry. ID will be provided only for 000 or 209 status. Otherwise, the ID will be null.
	InvoiceId	string	`json:"invoice_id"` //ID of the invoice related to the inquiry result.
	Timestamp	date.YYYYMMDDHHMMSS	`json:"timestamp"` //UTC Timestamp api hit (Format: yyyy-MM-ddTHH:mm:ss)
}

type AccountInquiryStatus struct {
	Code *string `json:"code"`
	Message *string `json:"message"`
}
