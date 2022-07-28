package scheduleddisbursement

import oygo "github.com/nadhirfr/oy-go"

type CreateParams struct {
	RecipientBank    oygo.BankCode `json:"recipient_bank"`    //Bank Code of the Beneficiary account, see Disbursement Bank Codes
	RecipientAccount string        `json:"recipient_account"` //Beneficiary account number, numeric only
	Amount           int           `json:"amount"`            //Amount of disbursement (Accept Non-Decimal Number), min amount 10000
	Note             string        `json:"note"`              //Add Note to the payout
	PartnerTrxId     string        `json:"partner_trx_id"`    //Unique Payout ID for a specific request, generated by partner
	Email            string        `json:"email"`             //Email for invoice notification (Optional). Email can be more than one but not more than five separated by a whitespace
	ScheduleDate     string        `json:"schedule_date"`     //Date for scheduled non-trigger-based disburse in "dd-mm-yyyy" format. Required if If is_trigger_based = FALSE or not set. Date is based on GMT+7
	IsTriggerBased   bool          `json:"is_trigger_based"`  //Whether scheduled transfer is trigger-based. Default value is false, if set to true, trigger_date and trigger_email are required
	TriggerDate      string        `json:"trigger_date"`      //Date when the disburse can be claimed by Beneficiary in "dd-mm-yyyy" format. Required if is_trigger_based = TRUE. Date is based on GMT+7
	TriggerEmail     string        `json:"trigger_email"`     //Email which the fund acceptance email and URL will be sent to. Required if is_trigger_based = TRUE. Only one email can be provided
	CsPhoneNumber    string        `json:"cs_phone_number"`   //Customer service phone number Beneficiary can contact when a trigger-based disbursement fails. This information will be shown to Beneficiary when trying to claim disbursement on a trigger-based disbursement and disburse execution fails due to insufficient balance. Required if is_trigger_based = TRUE
	CsEmail          string        `json:"cs_email"`          //Customer service email Beneficiary can contact when a trigger-based disbursement fails. This information will be shown to Beneficiary when trying to claim disbursement on a trigger-based disbursement and disburse execution fails due to insufficient balance. Required if is_trigger_based = TRUE
}

// GetParams contains parameters for Get
type GetParams struct {
	PartnerTxId string `json:"partner_tx_id" validate:"required"`
}