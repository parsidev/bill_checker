package billing

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_countDigits(t *testing.T) {
	type args struct {
		num *big.Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Pay ID",
			args: args{
				num: big.NewInt(int64(8283237)),
			},
			want: 7,
		},
		{
			name: "Bill ID",
			args: args{
				num: big.NewInt(int64(5152574430154)),
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.num); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitLastDigit(t *testing.T) {
	type args struct {
		number *big.Int
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 int
	}{
		{
			name: "Split Pay ID",
			args: args{
				number: big.NewInt(int64(8283237)),
			},
			want:  big.NewInt(int64(828323)),
			want1: 7,
		},
		{
			name: "Split Bill ID",
			args: args{
				number: big.NewInt(int64(5152574430154)),
			},
			want:  big.NewInt(int64(515257443015)),
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitLastDigit(tt.args.number)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("splitLastDigit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitLastDigit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_joinNumbers(t *testing.T) {
	res := new(big.Int)
	res.SetString("51525744301548283237", 10)

	type args struct {
		num1 *big.Int
		num2 *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test",
			args: args{
				num1: big.NewInt(int64(5152574430154)),
				num2: big.NewInt(int64(8283237)),
			},
			want: res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinNumbers(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joinNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkBill(t *testing.T) {
	type args struct {
		billID *big.Int
		payID  *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check bill",
			args: args{
				billID: big.NewInt(int64(5152574430154)),
				payID:  big.NewInt(int64(82832737)),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBill(tt.args.billID, tt.args.payID); got != tt.want {
				t.Errorf("checkBill() = %v, want %v", got, tt.want)
			}
		})
	}
}
