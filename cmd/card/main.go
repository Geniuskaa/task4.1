package main

import (
	"fmt"
	"github.com/Geniuskaa/task4.1/pkg/card"
	"github.com/Geniuskaa/task4.1/pkg/transfer"
)

func main() {
	service := card.NewService("Tinkoff Bank")
	visa := service.IssueCard("VISA","RUB")
	visa.Number = "4205 1348 6729 1672"
	visa.Balance = 1_350_34
	master := service.IssueCard("MasterCard", "RUB")
	master.Number = "7381 0375 2928 6823"
	master.Balance = 135_87
	//fmt.Println(service.Cards[0])
	fmt.Println(visa)
	fmt.Println(master)
	fmt.Println(service.Cards[0], service.Cards[1])

	transfer := transfer.NewService(service, 0.5, 10, 1.5, 30)
	total, ok := transfer.Card2Card("4205 1348 6729 1672", "7381 0375 2928 6823", 9_895_35)
	fmt.Println(total)
	fmt.Println(ok)
	fmt.Println(visa)
	total1, ok1 := transfer.Card2Card("7381 0375 2928 6823", "0381 01375 2920 6823", 5_895_35)
	fmt.Println(total1)
	fmt.Println(ok1)


}
