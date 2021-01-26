package delphi

import (
	"testing"
	"time"

	"github.com/solsw/gohelpers/mathhelper"
)

func TestTDateTimeToTime(t *testing.T) {
	type args struct {
		dt TDateTime
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// http://docwiki.embarcadero.com/Libraries/Sydney/en/System.TDateTime
		{name: "1", args: args{dt: 0}, want: "1899-12-30T00:00:00"},
		{name: "2", args: args{dt: 2.75}, want: "1900-01-01T18:00:00"},
		{name: "3", args: args{dt: 35065}, want: "1996-01-01T00:00:00"},
		{name: "4", args: args{dt: 35065.77}, want: "1996-01-01T18:28:48"},
		{name: "1n", args: args{dt: -1.25}, want: "1899-12-29T06:00:00"},
		{name: "2n", args: args{dt: -1.5}, want: "1899-12-29T12:00:00"},
		{name: "3n", args: args{dt: -1.9}, want: "1899-12-29T21:36:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TDateTimeToTime(tt.args.dt).In(time.UTC).Format("2006-01-02T15:04:05")
			if got != tt.want {
				t.Errorf("TDateTimeToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToTDateTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TDateTime
	}{
		{name: "1", args: args{t: time.Date(1899, time.Month(12), 30, 0, 0, 0, 0, time.UTC)}, want: 0},
		{name: "2", args: args{t: time.Date(1900, time.Month(1), 1, 18, 0, 0, 0, time.UTC)}, want: 2.75},
		{name: "3", args: args{t: time.Date(1996, time.Month(1), 1, 0, 0, 0, 0, time.UTC)}, want: 35065},
		{name: "4", args: args{t: time.Date(1996, time.Month(1), 1, 18, 28, 48, 0, time.UTC)}, want: 35065.77},
		{name: "1n", args: args{t: time.Date(1899, time.Month(12), 29, 6, 0, 0, 0, time.UTC)}, want: -1.25},
		{name: "2n", args: args{t: time.Date(1899, time.Month(12), 29, 12, 0, 0, 0, time.UTC)}, want: -1.5},
		{name: "3n", args: args{t: time.Date(1899, time.Month(12), 29, 21, 36, 0, 0, time.UTC)}, want: -1.9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToTDateTime(tt.args.t); !mathhelper.NearlyEqual(got, tt.want, 0.00001) {
				t.Errorf("TimeToTDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
