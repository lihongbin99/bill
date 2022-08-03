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

		start := false
		for i := len(records) - 1; i >= 0; i-- {
			record := records[i]
			if len(record) < 1 {
				continue
			}

			// 判断那里开始
			if !start {
				if strings.TrimSpace(record[0]) == "------------------------------------------------------------------------------------" {
					start = true
					continue
				}
			}

			// 没开始则跳过
			if !start {
				continue
			}

			// 判断哪里结束
			if strings.TrimSpace(record[0]) == "收/支" {
				break
			}

			if len(record) != 12 {
				return nil, fmt.Errorf("wechat bill len error(%d): %s", len(record), record)
			}

			transactionHour, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(record[10]))
			if err != nil {
				return nil, err
			}

			bill := model.Bill{
				App:                    "AliPay",
				TransactionHour:        transactionHour,
				TransactionType:        strings.TrimSpace(record[7]),
				Counterparty:           strings.TrimSpace(record[1]),
				OtherAccounts:          strings.TrimSpace(record[2]),
				Commodity:              strings.TrimSpace(record[3]),
				IncomeExpenditure:      strings.TrimSpace(record[0]),
				Amount:                 strings.TrimSpace(record[5]),
				PaymentMethod:          strings.TrimSpace(record[4]),
				CurrentState:           strings.TrimSpace(record[6]),
				TransactionNumber:      strings.TrimSpace(record[8]),
				MerchantTrackingNumber: strings.TrimSpace(record[9]),
				Remark:                 "",
			}
			billList = append(billList, bill)
		}
	}

	return billList, nil
}
