package stringhelper

import (
	"testing"
)

var (
	invalidString = "фывапр" + string([]byte("йцукен")[1:])
)

func TestIsEmptyOrWhite(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: ""}, want: true},
		{name: "2", args: args{s: "  "}, want: true},
		{name: "3", args: args{s: "\t"}, want: true},
		{name: "4", args: args{s: "\t \n"}, want: true},
		{name: "5", args: args{s: "qwerty"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyOrWhite(tt.args.s); got != tt.want {
				t.Errorf("IsEmptyOrWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDigital(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "01", args: args{s: ""}, want: false},
		{name: "02", args: args{s: "invalidString"}, want: false},
		{name: "1", args: args{s: "23"}, want: true},
		{name: "2", args: args{s: "2q3"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDigital(tt.args.s); got != tt.want {
				t.Errorf("IsDigital() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty string", args: args{s: ""}, want: false},
		{name: "invalidString", args: args{s: invalidString}, want: false},
		{name: "1", args: args{s: "3"}, want: true},
		{name: "2", args: args{s: "."}, want: false},
		{name: "3", args: args{s: "ы"}, want: false},
		{name: "4", args: args{s: "3.1415926535"}, want: true},
		{name: "5", args: args{s: "3.1415926535e-2"}, want: true},
		{name: "6", args: args{s: "3.1415926535E4"}, want: true},
		{name: "7", args: args{s: "3.1415926535E+4"}, want: true},
		{name: "8", args: args{s: "3.1415926535+4"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "QWERTY"}, want: true},
		{name: "2", args: args{s: "Qwerty"}, want: false},
		{name: "3", args: args{s: "qwertY"}, want: false},
		{name: "4", args: args{s: "asdfgh"}, want: false},
		{name: "5", args: args{s: "Б"}, want: true},
		{name: "6", args: args{s: "г"}, want: false},
		{name: "7", args: args{s: "01234"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpper(tt.args.s); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUpperRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{r: 'R'}, want: true},
		{name: "2", args: args{r: 'z'}, want: false},
		{name: "3", args: args{r: 'Б'}, want: true},
		{name: "4", args: args{r: 'г'}, want: false},
		{name: "5", args: args{r: '0'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpperRune(tt.args.r); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "QWERTY"}, want: false},
		{name: "2", args: args{s: "Qwerty"}, want: false},
		{name: "3", args: args{s: "qwertY"}, want: false},
		{name: "4", args: args{s: "asdfgh"}, want: true},
		{name: "5", args: args{s: "Б"}, want: false},
		{name: "6", args: args{s: "г"}, want: true},
		{name: "7", args: args{s: "01234"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLower(tt.args.s); got != tt.want {
				t.Errorf("IsLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLowerRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{r: 'R'}, want: false},
		{name: "2", args: args{r: 'z'}, want: true},
		{name: "3", args: args{r: 'Б'}, want: false},
		{name: "4", args: args{r: 'г'}, want: true},
		{name: "5", args: args{r: '0'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowerRune(tt.args.r); got != tt.want {
				t.Errorf("IsLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
