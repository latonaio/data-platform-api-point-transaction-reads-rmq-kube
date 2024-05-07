package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-point-transaction-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-point-transaction-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersBySenderReceiver":
			func() {
				header = c.HeadersBySenderReceiver(mtx, input, output, errs, log)
			}()
		case "HeadersBySender":
			func() {
				header = c.HeadersBySender(mtx, input, output, errs, log)
			}()
		case "HeadersByReceiver":
			func() {
				header = c.HeadersByReceiver(mtx, input, output, errs, log)
			}()
		case "HeadersByObject":
			func() {
				header = c.HeadersByObject(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:				header,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.PointTransaction = %d", input.Header.PointTransaction)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.PointTransaction ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySenderReceiver(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Sender = %d", input.Header.Sender)

	where := fmt.Sprintf("%s\nAND header.Receiver = %d", input.Header.Receiver)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Receiver ASC, header.Sender ASC, header.PointTransactionObjectType, header.PointTransactionObject, header.PointTransaction ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySender(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Sender = %d", input.Header.Sender)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Receiver ASC, header.Sender ASC, header.PointTransactionObjectType, header.PointTransactionObject, header.PointTransaction ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByReceiver(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Receiver = %d", input.Header.Receiver)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.Receiver ASC, header.Sender ASC, header.PointTransactionObjectType, header.PointTransactionObject, header.PointTransaction ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByObject(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.PointTransactionObject = %d", input.Header.PointTransactionObject)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsCancelled ASC, header.PointTransaction ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
