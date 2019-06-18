package timehelper

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTime0(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{name: "1", want: time0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time0(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time0() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{name: "1", args: args{month: 6}, want: time.June},
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

func TestFirstDayOfMonth(t *testing.T) {
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
		{name: "1e", args: args{month: 0}, wantErr: true},
		{name: "2e", args: args{year: 2019, month: 20}, wantErr: true},
		{name: "1", args: args{year: 2019, month: 4}, want: ymd(2019, time.April, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstDayOfMonth(tt.args.year, tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstDayOfMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstDayOfMonth() = %v, want %v", got, tt.want)
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
		{name: "1", args: args{year: 2018, month: time.February}, want: ymd(2018, time.February, 28)},
		{name: "2", args: args{year: 2018, month: time.December}, want: ymd(2018, time.December, 31)},
		{name: "3", args: args{year: 2020, month: time.February}, want: ymd(2020, time.February, 29)},
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
		{name: "1e", args: args{month: 0}, wantErr: true},
		{name: "2e", args: args{year: 2019, month: 20}, wantErr: true},
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

func TestDateYMD(t *testing.T) {
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
		{name: "1e", args: args{year: 2018, month: 40, day: 3}, want: ymd(2018, time.April, 2), wantErr: true},
		{name: "2e", args: args{year: 2018, month: 4, day: 31}, want: ymd(2018, time.April, 2), wantErr: true},
		{name: "1", args: args{year: 2018, month: 4, day: 2}, want: ymd(2018, time.April, 2), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateYMD(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateYMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !got.IsZero() {
					t.Errorf("DateYMD() = %v, want zero Time", got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateYMD() = %v, want %v", got, tt.want)
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
		{name: "1e", args: args{year: 2018, weekday: time.Weekday(-123)}, wantErr: true},
		{name: "2e", args: args{year: 2018, weekday: time.Weekday(123)}, wantErr: true},
		{name: "1", args: args{year: 2018, weekday: time.Monday}, want: ymd(2018, time.January, 1)},
		{name: "2", args: args{year: 2019, weekday: time.Tuesday}, want: ymd(2019, time.January, 1)},
		{name: "3", args: args{year: 2019, weekday: time.Monday}, want: ymd(2019, time.January, 7)},
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
		{name: "2", args: args{year: 2019, weekday: time.Tuesday}, args2: args{year: 2019, weekday: time.Tuesday}},
		{name: "3", args: args{year: 2019, weekday: time.Monday}, args2: args{year: 2019, weekday: time.Monday}},
		{name: "4", args: args{year: 2020, weekday: time.Friday}, args2: args{year: 2020, weekday: time.Friday}},
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

func TestWeekdaysInYear(t *testing.T) {
	type args struct {
		year    int
		weekday time.Weekday
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "1e", args: args{year: 2018, weekday: time.Weekday(-123)}, wantErr: true},
		{name: "2e", args: args{year: 2018, weekday: time.Weekday(123)}, wantErr: true},
		{name: "1", args: args{year: 2019, weekday: time.Monday}, want: 52},
		{name: "2", args: args{year: 2019, weekday: time.Tuesday}, want: 53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeekdaysInYear(tt.args.year, tt.args.weekday)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeekdaysInYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WeekdaysInYear() = %v, want %v", got, tt.want)
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

func TestPrevClosestWeekday(t *testing.T) {
	type args struct {
		t       time.Time
		weekday time.Weekday
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e", args: args{weekday: time.Weekday(-10)}, wantErr: true},
		{name: "2e", args: args{weekday: time.Weekday(10)}, wantErr: true},
		{name: "1", args: args{t: ymd(2019, time.April, 22), weekday: time.Monday}, want: ymd(2019, time.April, 22)},
		{name: "2", args: args{t: ymd(2019, time.April, 27), weekday: time.Monday}, want: ymd(2019, time.April, 22)},
		{name: "3", args: args{t: ymd(2019, time.April, 22), weekday: time.Wednesday}, want: ymd(2019, time.April, 17)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrevClosestWeekday(tt.args.t, tt.args.weekday)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrevClosestWeekday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrevClosestWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextClosestWeekday(t *testing.T) {
	type args struct {
		t       time.Time
		weekday time.Weekday
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "1e", args: args{weekday: time.Weekday(-10)}, wantErr: true},
		{name: "2e", args: args{weekday: time.Weekday(10)}, wantErr: true},
		{name: "1", args: args{t: ymd(2019, time.April, 22), weekday: time.Monday}, want: ymd(2019, time.April, 22)},
		{name: "2", args: args{t: ymd(2019, time.April, 30), weekday: time.Thursday}, want: ymd(2019, time.May, 2)},
		{name: "3", args: args{t: ymd(2019, time.April, 27), weekday: time.Thursday}, want: ymd(2019, time.May, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextClosestWeekday(tt.args.t, tt.args.weekday)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextClosestWeekday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextClosestWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}
