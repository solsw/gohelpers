package mathhelper

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestIsEven", args: args{i: 23}, want: false},
		{name: "TestIsEven", args: args{i: 2332}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.i); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundInt64(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{x: 2.1}, want: 2},
		{name: "2", args: args{x: 2.5}, want: 3},
		{name: "3", args: args{x: 2.7}, want: 3},
		{name: "4", args: args{x: -2.1}, want: -2},
		{name: "5", args: args{x: -2.5}, want: -3},
		{name: "6", args: args{x: -2.7}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundInt64(tt.args.x); got != tt.want {
				t.Errorf("RoundInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundToEvenInt64(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{x: 2.1}, want: 2},
		{name: "2", args: args{x: 2.5}, want: 2},
		{name: "3", args: args{x: 2.7}, want: 3},
		{name: "4", args: args{x: -2.1}, want: -2},
		{name: "5", args: args{x: -2.5}, want: -2},
		{name: "6", args: args{x: -2.7}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundToEvenInt64(tt.args.x); got != tt.want {
				t.Errorf("RoundToEvenInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncInt64(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{x: 2.1}, want: 2},
		{name: "2", args: args{x: 2.5}, want: 2},
		{name: "3", args: args{x: 2.7}, want: 2},
		{name: "4", args: args{x: -2.1}, want: -2},
		{name: "5", args: args{x: -2.5}, want: -2},
		{name: "6", args: args{x: -2.7}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncInt64(tt.args.x); got != tt.want {
				t.Errorf("TruncInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApproximatelyEquals(t *testing.T) {
	type args struct {
		value1    float64
		value2    float64
		tolerance float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{value1: 23, value2: 23.01, tolerance: 0.001}, want: false},
		{name: "2", args: args{value1: 23, value2: 23.0001, tolerance: 0.001}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApproximatelyEquals(tt.args.value1, tt.args.value2, tt.args.tolerance); got != tt.want {
				t.Errorf("ApproximatelyEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}
