package transfer

import (
	"github.com/Geniuskaa/task4.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc                   *card.Service
		InterestFromOurService    float64
		MinSumFromOurService      int64
		InterestNotFromOurService float64
		MinSumNotFromOurService   int64
	}

	testService := card.NewService("Test Bank")
	testService.Add(146_361_76, "4205 1348 6729 1672")
	testService.Add(361_76, "7505 1348 6729 1872")
	testService.Add(106_361_76, "4005 1388 6709 1473")
	testService.Add(1_76, "8905 1348 9229 1370")
	testService.Add(34_76, "8805 1348 9229 1370")
	testService.Add(12_76, "9905 1348 9229 1370")

	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
		wantErr error
	}{
		{"FirstTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"5106 2148 6729 1672", "7505 1348 6729 1872", 234_67}, 0, false, ErrInvalidCardNumber}, // 235_84,true,nil
		{"SecondTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"8805 1348 9229 1370", "4005 1388 6709 1473", 576_67}, 0, false, ErrInvalidCardNumber}, // 579_55
		{"ThirdTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"4005 1388 6709 1473", "7515 1348 6729 3872", 274_67}, 0, false, ErrInvalidCardNumber}, // 276_04, true
		{"ThourthTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"9905 1348 9229 1370", "4905 1348 7729 1872", 234_67}, 0, false, ErrInvalidCardNumber}, // 235_84
		{"FithTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"4105 1948 6729 1372", "5106 2148 6729 1672", 234_67}, 0, false, ErrInvalidCardNumber}, // 234_67
		{"SixthTest", fields{testService, 0.5, 10, 1.5, 30},
			args{"4215 1378 6729 1879", "7805 1338 6729 1172", 234_67}, 0, false, ErrInvalidCardNumber}, // 238_19
	}
	for _, tt := range tests {
		s := &Service{
			CardSvc:                   tt.fields.CardSvc,
			InterestFromOurService:    tt.fields.InterestFromOurService,
			MinSumFromOurService:      tt.fields.MinSumFromOurService,
			InterestNotFromOurService: tt.fields.InterestNotFromOurService,
			MinSumNotFromOurService:   tt.fields.MinSumNotFromOurService,
		}
		gotTotal, gotOk, gotErr := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotTotal != tt.wantTotal {
			t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
		}
		if gotOk != tt.wantOk {
			t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
		}
		if gotErr != tt.wantErr {
			t.Errorf("Card2Card() gotErr = %v, want %v", gotErr, tt.wantOk)
		}

	}

}
