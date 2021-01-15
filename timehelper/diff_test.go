package timehelper

import (
	"testing"
	"time"
)

func TestDaysDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1960, 9, 1)}, want: 70},
		{name: "2", args: args{t1: DateYMD(2021, 2, 1), t2: DateYMD(2021, 1, 1)}, want: -31},
		{name: "3", args: args{t1: DateYMD(1, 1, 1), t2: DateYMD(1960, 6, 23)}, want: 715683},
		{name: "4", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1959, 2, 8)}, want: -501},
		{name: "5", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1985, 8, 22)}, want: 9191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysDiff(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("DaysDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeeksDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1960, 9, 1)}, want: 10},
		{name: "2", args: args{t1: DateYMD(2021, 2, 1), t2: DateYMD(2021, 1, 1)}, want: -4},
		{name: "3", args: args{t1: DateYMD(1, 1, 1), t2: DateYMD(1960, 6, 23)}, want: 102240},
		{name: "4", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1959, 2, 8)}, want: -71},
		{name: "5", args: args{t1: DateYMD(1960, 6, 23), t2: DateYMD(1985, 8, 22)}, want: 1313},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeeksDiff(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("WeeksDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthsDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{t1: DateYMD(2021, 1, 1), t2: DateYMD(2021, 1, 31)}, want: 0},
		{name: "1", args: args{t1: DateYMD(2021, 1, 1), t2: DateYMD(2021, 2, 2)}, want: 1},
		{name: "2", args: args{t1: DateYMD(2021, 2, 2), t2: DateYMD(2021, 1, 1)}, want: -1},
		{name: "3", args: args{t1: DateYMD(2021, 2, 1), t2: DateYMD(2021, 1, 2)}, want: 0},
		{name: "4", args: args{t1: DateYMD(2020, 12, 1), t2: DateYMD(2021, 1, 31)}, want: 1},
		{name: "5", args: args{t1: DateYMD(2020, 12, 31), t2: DateYMD(2021, 1, 1)}, want: 0},
		{name: "6", args: args{t1: DateYMD(2021, 1, 31), t2: DateYMD(2020, 12, 1)}, want: -1},
		{name: "7", args: args{t1: DateYMD(2021, 1, 1), t2: DateYMD(2020, 12, 31)}, want: 0},
		{name: "8", args: args{t1: DateYMD(2018, 12, 31), t2: DateYMD(2021, 1, 1)}, want: 24},
		{name: "9", args: args{t1: DateYMD(2018, 12, 15), t2: DateYMD(2021, 1, 15)}, want: 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthsDiff(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("MonthsDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearsDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{t1: DateYMD(2021, 1, 1), t2: DateYMD(2021, 8, 31)}, want: 0},
		{name: "1", args: args{t1: DateYMD(2020, 2, 28), t2: DateYMD(2021, 2, 28)}, want: 1},
		{name: "2", args: args{t1: DateYMD(2020, 2, 29), t2: DateYMD(2021, 2, 28)}, want: 0},
		{name: "3", args: args{t1: DateYMD(2018, 12, 31), t2: DateYMD(2021, 1, 1)}, want: 2},
		{name: "4", args: args{t1: DateYMD(2018, 12, 15), t2: DateYMD(2021, 1, 15)}, want: 2},
		{name: "5", args: args{t1: DateYMD(2020, 2, 8), t2: DateYMD(1960, 6, 23)}, want: -59},
		{name: "6", args: args{t1: DateYMD(2020, 8, 22), t2: DateYMD(1960, 6, 23)}, want: -60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearsDiff(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("YearsDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
