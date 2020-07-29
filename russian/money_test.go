package russian

import "testing"

func TestNRubles(t *testing.T) {
	type args struct {
		number   int64
		showZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0t", args: args{number: 0, showZero: true}, want: "0 рублей"},
		{name: "0f", args: args{number: 0, showZero: false}, want: ""},
		{name: "1t", args: args{number: 10000002, showZero: true}, want: "10000002 рубля"},
		{name: "1f", args: args{number: 10000001, showZero: false}, want: "10000001 рубль"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NRubles(tt.args.number, tt.args.showZero); got != tt.want {
				t.Errorf("NRubles() = %v, want %v", got, tt.want)
			}
		})
	}
}
