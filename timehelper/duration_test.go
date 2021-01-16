package timehelper

import (
	"testing"
)

func TestFromWeeks(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "168h0m0s"},
		{name: "2", args: args{v: 1.0 / 7}, want: "24h0m0s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromWeeks(tt.args.v).String(); got != tt.want {
				t.Errorf("FromWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromDays(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "24h0m0s"},
		{name: "2", args: args{v: 2.2}, want: "52h48m0s"},
		{name: "3", args: args{v: 1.0 / 3}, want: "8h0m0s"},
		{name: "4", args: args{v: 0.01}, want: "14m24s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromDays(tt.args.v).String(); got != tt.want {
				t.Errorf("FromDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromHours(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "1h0m0s"},
		{name: "2", args: args{v: 1.5}, want: "1h30m0s"},
		{name: "3", args: args{v: 1.0 / 60}, want: "1m0s"},
		{name: "4", args: args{v: 1.0 / 3600}, want: "1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromHours(tt.args.v).String(); got != tt.want {
				t.Errorf("FromHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromMinutes(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "1m0s"},
		{name: "2", args: args{v: 1441 + 1.0/60}, want: "24h1m1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromMinutes(tt.args.v).String(); got != tt.want {
				t.Errorf("FromMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromSeconds(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "1s"},
		{name: "2", args: args{v: 1.0 / 11}, want: "90.90909ms"},
		{name: "3", args: args{v: 86401}, want: "24h0m1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSeconds(tt.args.v).String(); got != tt.want {
				t.Errorf("FromSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromMilliseconds(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "1ms"},
		{name: "2", args: args{v: 1.0 / 8}, want: "125µs"},
		{name: "3", args: args{v: 1.0 / 9}, want: "111.111µs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromMilliseconds(tt.args.v).String(); got != tt.want {
				t.Errorf("FromMilliseconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromMicroseconds(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "1µs"},
		{name: "2", args: args{v: 10000}, want: "10ms"},
		{name: "3", args: args{v: 1.0 / 6}, want: "166ns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromMicroseconds(tt.args.v).String(); got != tt.want {
				t.Errorf("FromMicroseconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromTicks(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 1.0}, want: "100ns"},
		{name: "2", args: args{v: 0.25}, want: "25ns"},
		{name: "3", args: args{v: 1.0 / 3}, want: "33ns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromTicks(tt.args.v).String(); got != tt.want {
				t.Errorf("FromTicks() = %v, want %v", got, tt.want)
			}
		})
	}
}
