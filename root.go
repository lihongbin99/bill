package main

import (
	"bill/alipay"
	"bill/common/db"
	"bill/model"
	"bill/wechat"
	"fmt"
	"os"
	"runtime"
	"sort"
)

func main() {
	totalBill := make(model.Bills, 0)

	alipayBill, err := alipay.Parse("C:\\Users\\Lee\\Documents\\bill\\alipay")
	if err != nil {
		exit("parse alipay bill error", err)
	}
	wechatBill, err := wechat.Parse("C:\\Users\\Lee\\Documents\\bill\\wechat")
	if err != nil {
		exit("parse wechat bill error", err)
	}

	if alipayBill != nil {
		totalBill = append(totalBill, alipayBill...)
		alipayBill = nil
	}
	if wechatBill != nil {
		totalBill = append(totalBill, wechatBill...)
		wechatBill = nil
	}

	if len(totalBill) == 0 {
		exit("not get bill", nil)
	}

	runtime.GC()

	sort.Sort(totalBill)

	for _, bill := range totalBill {
		if err = db.AddBill(bill); err != nil {
			exit("add bill to db error", err)
		}
	}
}

func exit(msg string, err error) {
	fmt.Println(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("input exit")
	buf := make([]byte, 1)
	_, _ = os.Stdin.Read(buf)
	os.Exit(0)
}
