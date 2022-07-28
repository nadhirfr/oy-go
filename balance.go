package oygo

type BalanceGet struct {
	Status             BalanceStatus `json:"status"`             //Status of Payout in Object {code: <status_code>, message: <status_message>}
	Balance            *float64      `json:"balance"`            //Remaining balance (Accept non fraction number)
	OverdraftBalance   *float64      `json:"overdraftBalance"`   //Remaining overdraft balance (Accept non fraction number)
	OverbookingBalance *float64      `json:"overbookingBalance"` //Remaining overbooking balance (Accept non fraction number)
	PendingBalance     *float64      `json:"pendingBalance"`     //The cumulative balance of your pending transactions.
	AvailableBalance   *float64      `json:"availableBalance"`   //The total cumulative money of Balance + Available Overdraft - Pending Balance that you can use for disbursement.
	Timestamp          *string       `json:"timestamp"`          //Execution time of Disbursement in OY! system ("dd-MM-yyyy HH:mm:ss").
}

type BalanceStatus struct {
	Code    *string `json:"code"`
	Message *string `json:"message"`
}
