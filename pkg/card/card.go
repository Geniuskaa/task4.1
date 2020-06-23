package card

type Service struct {
	BankName string
	Cards []*Card
}

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
	Icon string
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