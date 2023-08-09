package wechat

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
			if len(record) <= 1 {
				continue
			}
			if record[0] == "交易时间" {
				break
			}

			if len(record) != 11 {
				return nil, fmt.Errorf("wechat bill len error(%d): %s", len(record), record)
			}

			transactionHour, err := time.Parse("2006-01-02 15:04:05", record[0])
			if err != nil {
				return nil, err
			}
			amount := strings.TrimSpace(record[5])
			if strings.HasPrefix(amount, "¥") {
				amount = amount[2:]
			}

			bill := model.Bill{
				App:                    "WeChat",
				TransactionHour:        transactionHour,
				TransactionType:        record[1],
				Counterparty:           record[2],
				OtherAccounts:          "",
				Commodity:              record[3],
				IncomeExpenditure:      record[4],
				Amount:                 amount,
				PaymentMethod:          record[6],
				CurrentState:           record[7],
				TransactionNumber:      record[8],
				MerchantTrackingNumber: record[9],
				Remark:                 record[10],
			}
			billList = append(billList, bill)
		}
	}

	return billList, nil
}
