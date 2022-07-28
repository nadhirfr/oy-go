package oygo

type DisbursementResult struct {
	Status           DisbursementStatus `json:"status"`            //Status of Payout in Object {code: <status_code>, message: <status_message>}
	Amount           *float64           `json:"amount"`            //Amount of disbursement (Accept Non-Decimal Number)
	RecipientBank    BankCode           `json:"recipient_bank"`    //Bank Code of the Beneficiary account, see Disbursement Bank Codes
	RecipientAccount string             `json:"recipient_account"` //Beneficiary account number
	TrxId            string             `json:"trx_id"`            //Unique Payout ID from OY!. Partner can use this ID for settlement
	PartnerTrxId     string             `json:"partner_trx_id"`    //Unique Payout ID which partner put on the Request
	Timestamp        *string            `json:"timestamp"`         //Execution time of Disbursement in OY! system ("dd-MM-yyyy HH:mm:ss").
}

type DisbursementStatus struct {
	Code    *string `json:"code"`
	Message *string `json:"message"`
}

type DisbursementGet struct {
	Status              DisbursementStatus `json:"status"`                //Status of Payout in Object {code: <status_code>, message: <status_message>}
	TxStatusDescription string             `json:"tx_status_description"` //additional information of status code (e.g. FORCE CREDIT)
	Amount              float64            `json:"amount"`                //Amount of disbursement (Accept Non-Decimal Number)
	RecipientName       string             `json:"recipient_name"`        //Account holder name of Beneficiary account number
	RecipientBank       BankCode           `json:"recipient_bank"`        //Bank Code of the Beneficiary account, see Disbursement Bank Codes
	RecipientAccount    string             `json:"recipient_account"`     //Beneficiary account number
	TrxId               string             `json:"trx_id"`                //Unique Payout ID from OY!. Partner can use this ID for settlement
	PartnerTrxId        string             `json:"partner_trx_id"`        //Unique Payout ID which partner put on the Request, generated by partner
	Timestamp           string             `json:"timestamp"`             //Execution time of API remit status in OY! system ("dd-MM-yyyy HH:mm:ss")
	CreatedDate         string             `json:"created_date"`          //Executionn time of Disbures in OY! system ("dd-MM-yyyy HH:mm:ss")
	LastUpdatedDate     string             `json:"last_updated_date"`     //Latest status change of a disbursement. Example from 'Pending' to 'Success' ("dd-MM-yyyy HH:mm:ss")
}

type DisbursementCallback struct {
	Status              DisbursementStatus `json:"status" form:"status"`                               //Status of Payout in Object {code: <status_code>, message: <status_message>}
	TxStatusDescription string             `json:"tx_status_description" form:"tx_status_description"` //additional information of status code (e.g. FORCE CREDIT)
	Amount              float64            `json:"amount" form:"amount"`                               //Amount of disbursement (Accept Non-Decimal Number)
	RecipientName       string             `json:"recipient_name" form:"recipient_name"`               //Account holder name of Beneficiary account number
	RecipientBank       BankCode           `json:"recipient_bank" form:"recipient_bank"`               //Bank Code of the Beneficiary account, see Disbursement Bank Codes
	RecipientAccount    string             `json:"recipient_account" form:"recipient_account"`         //Beneficiary account number
	TrxId               string             `json:"trx_id" form:"trx_id"`                               //Unique Payout ID from OY!. Partner can use this ID for settlement
	PartnerTrxId        string             `json:"partner_trx_id" form:"partner_trx_id"`               //Unique Payout ID which partner put on the Request, generated by partner
	Timestamp           string             `json:"timestamp" form:"timestamp"`                         //Execution time of API remit status in OY! system ("dd-MM-yyyy HH:mm:ss")
	CreatedDate         string             `json:"created_date" form:"created_date"`                   //Executionn time of Disbures in OY! system ("dd-MM-yyyy HH:mm:ss")
	LastUpdatedDate     string             `json:"last_updated_date" form:"last_updated_date"`         //Latest status change of a disbursement. Example from 'Pending' to 'Success' ("dd-MM-yyyy HH:mm:ss")
}
