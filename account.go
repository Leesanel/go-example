package main

import (
	"fmt"
	"github.com/scylladb/termtables"
)

var (
	key int
	loop = true
	balance = 100.0
	money float64
	note string
	flag = false
	err error

)

const (
	detail = "明细"
	balances = "余额"
	payout = "收支"
	description = "说明"
)

func main() {

	t := termtables.CreateTable()
	t.AddHeaders(detail,balances,payout,description)

	for {
		fmt.Println("\n-------------收支记账-------------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Println("请选择（1-4）：")

		_, err = fmt.Scanln(&key)
		if err != nil {
			fmt.Println("输入有误!")
		}


		switch key {
		case 1:
			fmt.Println("-------------当前收支明细记录--------------")
			//当有收支时，显示收支详情
			if flag {
				//fmt.Println(Headline)
				fmt.Println(t.Render())
			} else {
				fmt.Println("暂时没有收支..")
			}

		case 2:
			fmt.Println("---------------登记收入--------------------")
			fmt.Println("本次收入金额：")
			//用户输入收入金额
			_, err = fmt.Scanln(&money)
			if err != nil {
				fmt.Println("输入有误!")
			}
			//用户余额便发生变化
			balance += money
			fmt.Println("本次收入说明：")
			//用户输入收入说明
			_, err = fmt.Scanln(&note)
			if err != nil {
				fmt.Println("输入有误!")
			}
			t.AddRow("income", balance, money, note)
			flag = true

		case 3:
			fmt.Println("---------------登记支出--------------------")
			fmt.Println("本次支出金额：")
			//用户输入支出金额
			_, err = fmt.Scanln(&money)
			if err != nil {
				fmt.Println("输入有误!")
			}
			//根据用户余额对支出的金额做出判断
			if money > balance {
				fmt.Println("你的余额不足，请重新输入...")
				break
			}
			//支出后用户的余额便发生变化
			balance -= money
			fmt.Println("本次支出说明：")
			_, err = fmt.Scanln(&note)
			if err != nil {
				fmt.Println("输入有误!")
			}
			t.AddRow("outcome", balance, money, note)
			flag = true

		case 4:
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for {
				_, err = fmt.Scanln(&choice)
				if err != nil {
					fmt.Println("输入有误!")
				}
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入y/n")
			}

			if choice == "y" {
				loop = false
			}

		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {
			break
		}

	}
	fmt.Println("你已经退出记账")
}
