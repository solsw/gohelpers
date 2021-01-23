package timehelper

import (
	"testing"
	"time"
)

func TestValidMonth(t *testing.T) {
	type args struct {
		month time.Month
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1e", args: args{month: time.Month(0)}, wantErr: true},
		{name: "2e", args: args{month: time.Month(23)}, wantErr: true},
		{name: "1", args: args{month: time.Month(2)}},
		{name: "2", args: args{month: time.June}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidMonth(tt.args.month); (err != nil) != tt.wantErr {
				t.Errorf("ValidMonth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidDate(t *testing.T) {
	type args struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1e", args: args{year: 2020, month: time.Month(0), day: 2}, wantErr: true},
		{name: "2e", args: args{year: 2020, month: time.Month(100), day: 2}, wantErr: true},
		{name: "3e", args: args{year: 2020, month: time.January, day: 222}, wantErr: true},
		{name: "4e", args: args{year: 2021, month: time.April, day: 31}, wantErr: true},
		{name: "5e", args: args{year: 2020, month: time.February, day: 30}, wantErr: true},
		{name: "1", args: args{year: 1960, month: time.June, day: 23}},
		{name: "2", args: args{year: 2020, month: time.February, day: 29}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidDate(tt.args.year, tt.args.month, tt.args.day); (err != nil) != tt.wantErr {
				t.Errorf("ValidDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidClockTime(t *testing.T) {
	type args struct {
		hour int
		min  int
		sec  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "0e", args: args{hour: -1, min: -2, sec: -3}, wantErr: true},
		{name: "1e", args: args{hour: 78, min: 34, sec: 56}, wantErr: true},
		{name: "2e", args: args{hour: 12, min: 78, sec: 56}, wantErr: true},
		{name: "3e", args: args{hour: 12, min: 34, sec: 60}, wantErr: true},
		{name: "0", args: args{hour: 0, min: 0, sec: 0}},
		{name: "1", args: args{hour: 12, min: 34, sec: 56}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidClockTime(tt.args.hour, tt.args.min, tt.args.sec); (err != nil) != tt.wantErr {
				t.Errorf("ValidClockTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidWeekday(t *testing.T) {
	type args struct {
		weekday time.Weekday
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1e", args: args{weekday: time.Weekday(-1)}, wantErr: true},
		{name: "2e", args: args{weekday: time.Weekday(8)}, wantErr: true},
		{name: "1", args: args{weekday: time.Weekday(0)}},
		{name: "2", args: args{weekday: time.Thursday}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidWeekday(tt.args.weekday); (err != nil) != tt.wantErr {
				t.Errorf("ValidWeekday() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
