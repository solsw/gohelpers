package russian

import "testing"

func TestNDays(t *testing.T) {
	type args struct {
		n        int64
		showZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{n: 0, showZero: false}, want: ""},
		{name: "01", args: args{n: 0, showZero: true}, want: "0 дней"},
		{name: "1", args: args{n: 1, showZero: true}, want: "1 день"},
		{name: "2", args: args{n: 2, showZero: true}, want: "2 дня"},
		{name: "3", args: args{n: 3, showZero: true}, want: "3 дня"},
		{name: "88", args: args{n: 88, showZero: true}, want: "88 дней"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NDays(tt.args.n, tt.args.showZero); got != tt.want {
				t.Errorf("NDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNInWordsDays(t *testing.T) {
	type args struct {
		n        int64
		showZero bool
		withZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{n: 0, showZero: false, withZero: false}, want: ""},
		{name: "01", args: args{n: 0, showZero: false, withZero: true}, want: ""},
		{name: "02", args: args{n: 0, showZero: true, withZero: false}, want: "ноль дней"},
		{name: "03", args: args{n: 0, showZero: true, withZero: true}, want: "ноль дней"},
		{name: "1", args: args{n: 1}, want: "один день"},
		{name: "2", args: args{n: 2}, want: "два дня"},
		{name: "10000", args: args{n: 10000}, want: "десять тысяч дней"},
		{name: "1000004", args: args{n: 1000004}, want: "один миллион четыре дня"},
		{name: "1000040", args: args{n: 1000040, withZero: true}, want: "один миллион ноль тысяч сорок дней"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NInWordsDays(tt.args.n, tt.args.showZero, tt.args.withZero); got != tt.want {
				t.Errorf("NInWordsDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNInWordsWeeks(t *testing.T) {
	type args struct {
		n        int64
		showZero bool
		withZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "01", args: args{n: 0, showZero: false, withZero: true}, want: ""},
		{name: "1", args: args{n: 1}, want: "одна неделя"},
		{name: "2", args: args{n: 2}, want: "две недели"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NInWordsWeeks(tt.args.n, tt.args.showZero, tt.args.withZero); got != tt.want {
				t.Errorf("NInWordsWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
