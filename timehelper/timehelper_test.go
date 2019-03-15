package timehelper

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestIntAsMonth(t *testing.T) {
	type args struct {
		month int
	}
	tests := []struct {
		name    string
		args    args
		want    time.Month
		wantErr bool
	}{
		{name: "1e", args: args{month: 0}, wantErr: true},
		{name: "2e", args: args{month: -23}, wantErr: true},
		{name: "3e", args: args{month: 23}, wantErr: true},
		{name: "1", args: args{month: 6}, want: time.June, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntAsMonth(tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntAsMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntAsMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestLastDayOfMonth(t *testing.T) {
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e", args: args{year: 2018, month: time.Month(0)}, wantErr: true},
		{name: "1", args: args{year: 2018, month: time.February}, want: dateYMDPrim(2018, time.February, 28), wantErr: false},
		{name: "2", args: args{year: 2018, month: time.December}, want: dateYMDPrim(2018, time.December, 31), wantErr: false},
		{name: "3", args: args{year: 2020, month: time.February}, want: dateYMDPrim(2020, time.February, 29), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastDayOfMonth(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastDayOfMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastDayOfMonth() = %v, want %v", got, tt.want)
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
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "1", args: args{year: 2018, month: time.February}, want: 28},
		{name: "2", args: args{year: 2020, month: time.February}, want: 29},
		{name: "3", args: args{year: 2000, month: time.February}, want: 29},
		{name: "1", args: args{year: 2018, month: time.June}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysInMonth(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("DaysInMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DaysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateYMDEr(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e", args: args{year: 2018, month: 40, day: 3}, want: dateYMDPrim(2018, time.April, 2), wantErr: true},
		{name: "2e", args: args{year: 2018, month: 4, day: 31}, want: dateYMDPrim(2018, time.April, 2), wantErr: true},
		{name: "1", args: args{year: 2018, month: 4, day: 2}, want: dateYMDPrim(2018, time.April, 2), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateYMDEr(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateYMDEr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !got.IsZero() {
					t.Errorf("DateYMDEr() = %v, want zero Time", got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateYMDEr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstWeekdayInYear(t *testing.T) {
	type args struct {
		year    int
		weekday time.Weekday
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e", args: args{year: 2018, weekday: time.Weekday(-123)}, want: time0, wantErr: true},
		{name: "2e", args: args{year: 2018, weekday: time.Weekday(123)}, want: time0, wantErr: true},
		{name: "1", args: args{year: 2018, weekday: time.Monday}, want: dateYMDPrim(2018, time.January, 1), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstWeekdayInYear(tt.args.year, tt.args.weekday)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstWeekdayInYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstWeekdayInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoFirstWeekdayInYearEqual(t *testing.T) {
	type args struct {
		year    int
		weekday time.Weekday
	}
	tests := []struct {
		name  string
		args  args
		args2 args
	}{
		{name: "1e", args: args{year: 2018, weekday: time.Weekday(-123)}, args2: args{year: 2018, weekday: time.Weekday(123)}},
		{name: "1", args: args{year: 2018, weekday: time.Monday}, args2: args{year: 2018, weekday: time.Monday}},
		{name: "2", args: args{year: 2020, weekday: time.Friday}, args2: args{year: 2020, weekday: time.Friday}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstWeekdayInYear(tt.args.year, tt.args.weekday)
			got2, err2 := FirstWeekdayInYear2(tt.args2.year, tt.args2.weekday)
			if !reflect.DeepEqual(err, err2) {
				t.Errorf("FirstWeekdayInYear() error = %v, FirstWeekdayInYear2() error = %v", err, err2)
				return
			}
			if got != got2 {
				t.Errorf("FirstWeekdayInYear() = %v, FirstWeekdayInYear2() = %v", got, got2)
			}
		})
	}
}

func TestTwoFridaysInYearEqual(t *testing.T) {
	type args struct {
		year int
	}
	type test struct {
		name string
		args args
	}
	tests := make([]test, 0)
	tests = append(tests, test{name: "1800", args: args{year: 1800}})
	for i := 0; i < 10; i++ {
		tests = append(tests, test{name: fmt.Sprint(i), args: args{year: 2000 + i}})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FridaysInYear(tt.args.year)
			got2 := FridaysInYear2(tt.args.year)
			if got != got2 {
				t.Errorf("FridaysInYear() = %v, FridaysInYear2() = %v", got, got2)
			}
		})
	}
}
