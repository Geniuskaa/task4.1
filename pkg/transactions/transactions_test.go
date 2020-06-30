package transactions

import (
	"github.com/Geniuskaa/task4.1/pkg/card"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {

	testService := card.NewService("Test Bank")
	visa := testService.IssueCard("VISA","RUB")
	masterCard := testService.IssueCard("MasterCard","RUB")
	visaMirror := testService.IssueCard("VISA","RUB")
	masterCardMirror := testService.IssueCard("MasterCard","RUB")

	visa.Transactions = make([]*card.Transaction, 4)
	masterCard.Transactions = make([]*card.Transaction, 3)

	visaMirror.Transactions = make([]*card.Transaction, 4)
	masterCardMirror.Transactions = make([]*card.Transaction, 3)


	transaction10 := card.Transaction{
		Id: 10,
		SumOfTransaction: 146_098_00,
		Date: 1592657460,
		MCC: "5920",
		Status: "Обработано",
	}
	tr10 := &transaction10

	transaction20 := card.Transaction{
		Id: 20,
		SumOfTransaction: 6_198_00,
		Date: 1592657760,
		MCC: "5920",
		Status: "Обработано",
	}
	tr20 := &transaction20

	transaction30 := card.Transaction{
		Id: 30,
		SumOfTransaction: 19_498_00,
		Date: 1592657360,
		MCC: "5920",
		Status: "Обработано",
	}
	tr30 := &transaction30

	transaction40 := card.Transaction{
		Id: 40,
		SumOfTransaction: 10_498_00,
		Date: 1592657060,
		MCC: "5920",
		Status: "Обработано",
	}
	tr40 := &transaction40

	transaction110 := card.Transaction{
		Id: 60,
		SumOfTransaction: 12_498_00,
		Date: 1592657060,
		MCC: "5920",
		Status: "Обработано",
	}
	tr110 := &transaction110

	visa.Transactions[0] = tr10
	visa.Transactions[1] = tr20
	visa.Transactions[2] = tr30
	visa.Transactions[3] = tr40

	visaMirror.Transactions[0] = tr10
	visaMirror.Transactions[1] = tr30
	visaMirror.Transactions[2] = tr40
	visaMirror.Transactions[3] = tr20

	masterCard.Transactions[0] = tr40
	masterCard.Transactions[1] = tr110
	masterCard.Transactions[2] = tr20

	masterCardMirror.Transactions[0] = tr110
	masterCardMirror.Transactions[1] = tr40
	masterCardMirror.Transactions[2] = tr20

	type args struct {
		slice []*card.Transaction
	}

	tests := []struct {
		name string
		args args
		wantArg []*card.Transaction
	}{
		{"transaction1", args{visa.Transactions}, visaMirror.Transactions },
		{"transaction2", args{masterCard.Transactions}, masterCardMirror.Transactions},
	}

	for _, tt := range tests {
		if got := Sort(tt.args.slice); !reflect.DeepEqual(tt.args.slice, tt.wantArg) {
			t.Errorf("Sum() = %v, want %v", got, tt.wantArg)
		}
	}
}