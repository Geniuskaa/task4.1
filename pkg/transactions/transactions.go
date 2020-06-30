package transactions

import (
	"github.com/Geniuskaa/task4.1/pkg/card"
	"sort"
)


type Service struct {
	CardTransaction *card.Service
}

func AddNewTransaction(card *card.Card, transaction card.Transaction) *card.Card {
	card.Transactions = append(card.Transactions, &transaction)
	return card
}

func Sort(slice []*card.Transaction) []*card.Transaction{
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].SumOfTransaction > slice[j].SumOfTransaction })
	return slice
}