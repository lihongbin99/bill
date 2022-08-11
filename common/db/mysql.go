package db

import (
	"bill/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	d, err := sql.Open("mysql", "root:root@tcp(192.168.42.148:3306)/my_data?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db = d
}

func AddBill(bill model.Bill) error {
	_, err := db.Exec("INSERT INTO bill VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		nil,
		bill.App,
		bill.TransactionHour,
		bill.TransactionType,
		bill.Counterparty,
		bill.OtherAccounts,
		bill.Commodity,
		bill.IncomeExpenditure,
		bill.Amount,
		bill.PaymentMethod,
		bill.CurrentState,
		bill.TransactionNumber,
		bill.MerchantTrackingNumber,
		bill.Remark,
	)

	return err
}
