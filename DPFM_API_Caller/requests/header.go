package requests

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
