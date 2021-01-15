package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestWeekdaysInYear(t *testing.T) {
	type args struct {
		year    int
		weekday time.Weekday
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{year: 2019, weekday: time.Monday}, want: 52},
		{name: "2", args: args{year: 2019, weekday: time.Tuesday}, want: 53},
		{name: "3", args: args{year: 2020, weekday: time.Wednesday}, want: 53},
		{name: "4", args: args{year: 2020, weekday: time.Thursday}, want: 53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekdaysInYear(tt.args.year, tt.args.weekday); got != tt.want {
				t.Errorf("WeekdaysInYear() = %v, want %v", got, tt.want)
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
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{t: DateYMD(2019, time.April, 22), weekday: time.Monday}, want: DateYMD(2019, time.April, 22)},
		{name: "2", args: args{t: DateYMD(2019, time.April, 27), weekday: time.Monday}, want: DateYMD(2019, time.April, 22)},
		{name: "3", args: args{t: DateYMD(2019, time.April, 22), weekday: time.Wednesday}, want: DateYMD(2019, time.April, 17)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrevClosestWeekday(tt.args.t, tt.args.weekday); !reflect.DeepEqual(got, tt.want) {
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
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{t: DateYMD(2019, time.April, 22), weekday: time.Monday}, want: DateYMD(2019, time.April, 22)},
		{name: "2", args: args{t: DateYMD(2019, time.April, 30), weekday: time.Thursday}, want: DateYMD(2019, time.May, 2)},
		{name: "3", args: args{t: DateYMD(2019, time.April, 27), weekday: time.Thursday}, want: DateYMD(2019, time.May, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextClosestWeekday(tt.args.t, tt.args.weekday); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextClosestWeekday() = %v, want %v", got, tt.want)
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
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{year: 2018, weekday: time.Monday}, want: DateYMD(2018, time.January, 1)},
		{name: "2", args: args{year: 2019, weekday: time.Tuesday}, want: DateYMD(2019, time.January, 1)},
		{name: "3", args: args{year: 2019, weekday: time.Monday}, want: DateYMD(2019, time.January, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstWeekdayInYear(tt.args.year, tt.args.weekday); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstWeekdayInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastWeekdayInYear(t *testing.T) {
	type args struct {
		year    int
		weekday time.Weekday
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{year: 2020, weekday: time.Thursday}, want: DateYMD(2020, time.December, 31)},
		{name: "2", args: args{year: 2020, weekday: time.Wednesday}, want: DateYMD(2020, time.December, 30)},
		{name: "3", args: args{year: 2020, weekday: time.Saturday}, want: DateYMD(2020, time.December, 26)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastWeekdayInYear(tt.args.year, tt.args.weekday); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastWeekdayInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
