package alipay

import (
	"bill/common/utils"
	"bill/model"
	csv2 "encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func Parse(csvPath string) ([]model.Bill, error) {
	billList := make([]model.Bill, 0)

	files, err := utils.GetAllFile(csvPath)
	if err != nil || files == nil {
		return nil, err
	}

	for _, fileName := range files {
		filePath := fmt.Sprintf("%s%c%s", csvPath, os.PathSeparator, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		csvReader := csv2.NewReader(file)
		csvReader.Comma = ','
		csvReader.FieldsPerRecord = -1

		records, err := csvReader.ReadAll()
		if err != nil {
			return nil, err
		}

		for i := len(records) - 1; i >= 0; i-- {
			record := records[i]
			if len(record) < 1 {
				continue
			}
			if record[0] == "交易时间" {
				break
			}

			if len(record) != 13 {
				return nil, fmt.Errorf("wechat bill len error(%d): %s", len(record), record)
			}

			transactionHour, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(record[0]))
			if err != nil {
				return nil, err
			}

			bill := model.Bill{
				App:                    "AliPay",
				TransactionHour:        transactionHour,
				TransactionType:        strings.TrimSpace(record[1]),
				Counterparty:           strings.TrimSpace(record[2]),
				OtherAccounts:          strings.TrimSpace(record[3]),
				Commodity:              strings.TrimSpace(record[4]),
				IncomeExpenditure:      strings.TrimSpace(record[5]),
				Amount:                 strings.TrimSpace(record[6]),
				PaymentMethod:          strings.TrimSpace(record[7]),
				CurrentState:           strings.TrimSpace(record[8]),
				TransactionNumber:      strings.TrimSpace(record[9]),
				MerchantTrackingNumber: strings.TrimSpace(record[10]),
				Remark:                 strings.TrimSpace(record[11]),
			}
			billList = append(billList, bill)
		}
	}

	return billList, nil
}
