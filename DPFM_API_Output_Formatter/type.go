package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header	*[]Header	`json:"Header"`
}

type Header struct {
	PointTransaction						int		`json:"PointTransaction"`
	PointTransactionType					string	`json:"PointTransactionType"`
	PointTransactionDate					string	`json:"PointTransactionDate"`
	PointTransactionTime					string	`json:"PointTransactionTime"`
	Sender									int		`json:"Sender"`
	Receiver								int		`json:"Receiver"`
	PointSymbol								string	`json:"PointSymbol"`
	PlusMinus								string	`json:"PlusMinus"`
	PointTransactionAmount					float32	`json:"PointTransactionAmount"`
	PointTransactionObjectType				string	`json:"PointTransactionObjectType"`
	PointTransactionObject					int		`json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction		float32	`json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction		float32	`json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction	float32	`json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction	float32	`json:"ReceiverPointBalanceAfterTransaction"`
	Attendance								*int	`json:"Attendance"`
	Participation							*int	`json:"Participation"`
	CreationDate							string	`json:"CreationDate"`
	CreationTime							string	`json:"CreationTime"`
	IsCancelled								*bool	`json:"IsCancelled"`
}
