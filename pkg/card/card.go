package card

import (
	"strconv"
	"strings"
)

type Service struct {
	BankName string
	Cards []*Card
}

type Transaction struct {
	Id               int64
	SumOfTransaction int64
	Date             int64 // в формате Unix Timeship
	MCC              string
	Status           string
}

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
	Icon string
	Transactions []*Transaction
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(issuer string, currency string) *Card {
	card := &Card{
		Issuer: issuer,
		Balance: 0,
		Currency: currency,
		Number: "0000 0000 0000 0000",
		Icon: "https://.......",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) Add(balance int64, number string) *Card{
	card := &Card{
		Issuer: "issuer",
		Balance: balance,
		Currency: "RUB",
		Number: number,
		Icon: "https://.......",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) Find(cardNumber string) *Card {
	for _, card := range s.Cards {
		if card.Number == cardNumber {
			return card
		}
	}
	return nil
}

func IsValid(number string) (valid bool)  {
	n := strings.Split(strings.ReplaceAll(number, " ", ""), "")
	a := make([]int, len(n))
	i := 0
	for _, m := range n {
		a[i], _ = strconv.Atoi(m)
		i++
	}

	sum := 0
	for i=0; i < len(a); i++ {
		if i == 0 || i % 2 == 0 {
			if a[i]*2 > 9 {
				sum = a[i]*2 - 9
			}
			sum += a[i]
		}
		sum += a[i]
	}
	if sum % 10 != 0 {
		return false
	}
	return true
}

