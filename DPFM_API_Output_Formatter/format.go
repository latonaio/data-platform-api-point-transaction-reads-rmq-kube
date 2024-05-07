package dpfm_api_output_formatter

import (
	"data-platform-api-point-transaction-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.PointTransaction,
			&pm.PointTransactionType,
			&pm.PointTransactionDate,
			&pm.PointTransactionTime,
			&pm.Sender,
			&pm.Receiver,
			&pm.PointSymbol,
			&pm.PlusMinus,
			&pm.PointAmount,
			&pm.PointTransactionObjectType,
			&pm.PointTransactionObject,
			&pm.SenderPointBalanceBeforeTransaction,
			&pm.SenderPointBalanceAfterTransaction,
			&pm.ReceiverPointBalanceBeforeTransaction,
			&pm.ReceiverPointBalanceAfterTransaction,
			&pm.Attendance,
			&pm.Participation,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			PointTransaction:						data.PointTransaction,
			PointTransactionType:					data.PointTransactionType,
			PointTransactionDate:					data.PointTransactionDate,
			PointTransactionTime:					data.PointTransactionTime,
			Sender:									data.Sender,
			Receiver:								data.Receiver,
			PointSymbol:							data.PointSymbol,
			PlusMinus:								data.PlusMinus,
			PointAmount:							data.PointAmount,
			PointTransactionObjectType:				data.PointTransactionObjectType,
			PointTransactionObject:					data.PointTransactionObject,
			SenderPointBalanceBeforeTransaction:	data.SenderPointBalanceBeforeTransaction,
			SenderPointBalanceAfterTransaction:		data.SenderPointBalanceAfterTransaction,
			ReceiverPointBalanceBeforeTransaction:	data.ReceiverPointBalanceBeforeTransaction,
			ReceiverPointBalanceAfterTransaction:	data.ReceiverPointBalanceAfterTransaction,
			Attendance:								data.Attendance,
			Participation:							data.Participation,
			CreationDate:							data.CreationDate,
			CreationTime:							data.CreationTime,
			IsCancelled:							data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
