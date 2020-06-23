package transfer

import (
	"fmt"
	"github.com/Geniuskaa/task4.1/pkg/card"
)

type Service struct {
	CardSvc                   *card.Service
	InterestFromOurService    float64 // 0.5% + min 10 RUB (FromOurService)
	MinSumFromOurService      int64
	InterestNotFromOurService float64 // 1.5% + min 30 RUB
	MinSumNotFromOurService   int64
}

func NewService(cardSvc *card.Service, interestFOS float64, minSumFOS int64, interestNFOS float64, minSiumNFOS int64) *Service {
	return &Service{CardSvc: cardSvc, InterestFromOurService: interestFOS, MinSumFromOurService: minSumFOS,
		InterestNotFromOurService: interestNFOS, MinSumNotFromOurService: minSiumNFOS}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) { //"6373 0285 2950 1052"
	f := false
	fIndex := 0
	t := false
	tIndex := 0
	i := 0
	for i < len(s.CardSvc.Cards) {
		if from == s.CardSvc.Cards[i].Number {
			f = true
			fIndex = i
			break
		} else { i++ }
	}

	i = 0

	for i < len(s.CardSvc.Cards) {
		if to == s.CardSvc.Cards[i].Number {
			t = true
			tIndex = i
			break
		} else { i++ }
	}

	i = 0

	switch {
	case f == true && t == true: //fromOur
		if s.MinSumFromOurService > amount {
			fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
			return 0, false
		} else {
			sum := float64(amount) * (s.InterestFromOurService/100.0 + 1.0)
			total := int64(sum)
			if s.CardSvc.Cards[fIndex].Balance < total {
				fmt.Println("Недостачно денег на балансе для перевода")
				return total, false
			} else {
				s.CardSvc.Cards[fIndex].Balance -= total
				s.CardSvc.Cards[tIndex].Balance += amount
			}
			return total, true
		}
	case f == false && t == false: //fromNotOur
		if s.MinSumNotFromOurService > amount {
			fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
			return 0, false
		} else {
			sum := float64(amount) * (s.InterestNotFromOurService/100.0 + 1.0)
			total := int64(sum)
			return total, true
		}
	case f == true && t == false: //fromOur
		if s.MinSumFromOurService > amount {
			fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
			return 0, false
		} else {
			sum := float64(amount) * (s.InterestFromOurService/100.0 + 1.0)
			total := int64(sum)
			if s.CardSvc.Cards[fIndex].Balance < total {
				fmt.Println("Недостачно денег на балансе для перевода")
				return total, false
			} else {
				s.CardSvc.Cards[fIndex].Balance -= total
			}
			return total, true
		}
	case f == false && t == true: //NoInterest

		/*if s.MinSumFromOurService > amount {
			fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
		} else
			sum := float64(amount) * (s.InterestNotFromOurService/100.0 + 1.0)*/

		total := amount //int64(sum)
		s.CardSvc.Cards[tIndex].Balance += total
		return total, true
	}
	return 0, true
}