package model

import (
	"encoding/json"
	"time"
)

type Bill struct {
	App                    string
	TransactionHour        time.Time // AliPay/WeChat 交易时间
	TransactionType        string    // AliPay/WeChat 交易类型
	Counterparty           string    // AliPay/WeChat 交易对方
	OtherAccounts          string    // AliPay        对方账号
	Commodity              string    // AliPay/WeChat 商品
	IncomeExpenditure      string    // AliPay/WeChat 收/支
	Amount                 string    // AliPay/WeChat 金额
	PaymentMethod          string    // AliPay/WeChat 支付方式
	CurrentState           string    // AliPay/WeChat 当前状态
	TransactionNumber      string    // AliPay/WeChat 交易单号
	MerchantTrackingNumber string    // AliPay/WeChat 商户单号
	Remark                 string    //       /WeChat 备注
}

func (that *Bill) String() string {
	bytes, err := json.Marshal(that)
	if err != nil {
		return "\"error\":\"" + err.Error() + "\""
	}
	return string(bytes)
}

type Bills []Bill

func (that Bills) Len() int {
	return len(that)
}

func (that Bills) Less(i, j int) bool {
	return that[i].TransactionHour.UnixMilli() < that[j].TransactionHour.UnixMilli()
}

func (that Bills) Swap(i, j int) {
	that[i], that[j] = that[j], that[i]
}
