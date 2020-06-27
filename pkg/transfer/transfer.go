package transfer

import (
	"errors"
	"github.com/Geniuskaa/task4.1/pkg/card"
)

var (
	ErrCardBalanceNotEnough = errors.New("Not enough money on the balance")
	ErrAmountOfTransferTooLow = errors.New("The sum of transfer is less than the minimum amount")
	ErrUnknowReason = errors.New("Something bad happend. Try again later")
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

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool, err error) { //"6373 0285 2950 1052"
	fInBase := s.CardSvc.Find(from)
	tInBase := s.CardSvc.Find(to)

	switch {
	case fInBase != nil && tInBase != nil: //fromOur
		if s.MinSumFromOurService > amount {
			return 0, false, ErrAmountOfTransferTooLow //fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
		} else {
			sum := float64(amount) * (s.InterestFromOurService/100.0 + 1.0)
			total := int64(sum)
			if fInBase.Balance < total {
				return total, false, ErrCardBalanceNotEnough  // fmt.Println("Недостачно денег на балансе для перевода")
			}
			fInBase.Balance -= total
			tInBase.Balance += amount

			return total, true, nil
		}
	case fInBase == nil && tInBase == nil: //fromNotOur
		if s.MinSumNotFromOurService > amount {
			return 0, false, ErrAmountOfTransferTooLow  // fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
		} else {
			sum := float64(amount) * (s.InterestNotFromOurService/100.0 + 1.0)
			total := int64(sum)
			return total, true, nil
		}
	case fInBase != nil && tInBase == nil: //fromOur
		if s.MinSumFromOurService > amount {
			return 0, false, ErrAmountOfTransferTooLow  // fmt.Println("Сумма желаемого перевода меньше минимальной суммы перевода")
		} else {
			sum := float64(amount) * (s.InterestFromOurService/100.0 + 1.0)
			total := int64(sum)
			if fInBase.Balance < total {
				return total, false, ErrCardBalanceNotEnough  // fmt.Println("Недостачно денег на балансе для перевода")
			}
				fInBase.Balance -= total

			return total, true, nil
		}
	case fInBase == nil && tInBase != nil: //NoInterest
		total := amount //int64(sum)
		tInBase.Balance += total
		return total, true, nil
	default:
		return 0, false, ErrUnknowReason
	}
	return 0, false, ErrUnknowReason
}

/*type TransferError struct {
	message string
}

func (e *TransferError) Error() string {
	return e.message
}*/
