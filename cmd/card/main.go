package main

import (
	"fmt"
	"github.com/Geniuskaa/task4.1/pkg/card"
	"github.com/Geniuskaa/task4.1/pkg/transactions"
)

func main() {
	service := card.NewService("Tinkoff Bank")

	visa := service.IssueCard("VISA","RUB")
	visa.Number = "5106 2148 6729 1672"
	visa.Balance = 1_350_34

	master := service.IssueCard("MasterCard", "RUB")
	master.Number = "7381 0375 2928 6823"
	master.Balance = 135_87


	fmt.Println(visa)
	fmt.Println(master)
	fmt.Println(service.Cards[0], service.Cards[1])

	/*transfer := transfer.NewService(service, 0.5, 10, 1.5, 30)
	total, ok, err:= transfer.Card2Card("4205 1348 6729 1672", "7381 0375 2928 6823", 9_895_35)
	fmt.Println(total)
	fmt.Println(ok)
	fmt.Println(err)
	fmt.Println(transfer.Card2Card("4205 1348 6729 1672", "7381 0375 2928 6823", 9_895_35))
	fmt.Println(visa)
	total1, ok1, err := transfer.Card2Card("7381 0375 2928 6823", "0381 01375 2920 6823", 5_895_35)
	fmt.Println(total1)
	fmt.Println(ok1)*/
	//visa = transactions.AddNewTransaction(transaction, )
	transaction1 := card.Transaction{
		Id: 1,
		SumOfTransaction: 146_498_00,
		Date: 1592657460,
		MCC: "5920",
		Status: "Обработано",
	}

	transaction2 := card.Transaction{
		Id: 2,
		SumOfTransaction: 6_498_00,
		Date: 1592657760,
		MCC: "5420",
		Status: "Обработано",
	}

	transaction3 := card.Transaction{
		Id: 3,
		SumOfTransaction: 1_498_00,
		Date: 1592657360,
		MCC: "5920",
		Status: "Обработано",
	}

	transaction4 := card.Transaction{
		Id: 4,
		SumOfTransaction: 16_498_00,
		Date: 1592657060,
		MCC: "5420",
		Status: "Обработано",
	}
	visa = transactions.AddNewTransaction(visa, transaction1)
	visa = transactions.AddNewTransaction(visa, transaction2)
	visa = transactions.AddNewTransaction(visa, transaction3)
	visa = transactions.AddNewTransaction(visa, transaction4)


	transactions.Sort(visa.Transactions)

	for i := 0; i < len(visa.Transactions); i++ {
			fmt.Println(visa.Transactions[i])
		}





}
