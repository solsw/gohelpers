package timehelper

import (
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{year: 1800}, want: false},
		{name: "2", args: args{year: 1900}, want: false},
		{name: "3", args: args{year: 2000}, want: true},
		{name: "4", args: args{year: 2018}, want: false},
		{name: "5", args: args{year: 2020}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0e", args: args{year: 1960, month: time.Month(0)}, want: -1},
		{name: "1e", args: args{year: 1960, month: time.Month(100)}, want: -1},
		{name: "0", args: args{year: 1960, month: time.June}, want: 30},
		{name: "1", args: args{year: 2018, month: time.February}, want: 28},
		{name: "2", args: args{year: 2020, month: time.February}, want: 29},
		{name: "3", args: args{year: 2000, month: time.February}, want: 29},
		{name: "4", args: args{year: 1900, month: time.Month(2)}, want: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysInMonth(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("DaysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
