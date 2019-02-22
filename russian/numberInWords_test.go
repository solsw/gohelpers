package russian

import "testing"

func TestNumberInWords(t *testing.T) {
	type args struct {
		number    int64
		gender    GrammaticalGender
		withZeros bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0mt", args: args{number: 0, gender: Masculine, withZeros: true}, want: "ноль"},
		{name: "0mf", args: args{number: 0, gender: Masculine, withZeros: false}, want: "ноль"},
		{name: "-1mt", args: args{number: -1, gender: Masculine, withZeros: true}, want: "минус один"},
		{name: "1ft", args: args{number: 1, gender: Feminine, withZeros: true}, want: "одна"},
		{name: "100ff", args: args{number: 100, gender: Feminine, withZeros: false}, want: "сто"},
		{name: "1000mt", args: args{number: 1000, gender: Masculine, withZeros: true}, want: "одна тысяча ноль"},
		{name: "1000mf", args: args{number: 1000, gender: Masculine, withZeros: false}, want: "одна тысяча"},
		{name: "1001mf", args: args{number: 1001, gender: Masculine, withZeros: false}, want: "одна тысяча один"},
		{name: "2000002ft", args: args{number: 2000002, gender: Feminine, withZeros: true}, want: "два миллиона ноль тысяч две"},
		{name: "2000002ff", args: args{number: 2000002, gender: Feminine, withZeros: false}, want: "два миллиона две"},
		{name: "2000000000000002ff", args: args{number: 2000000000000002, gender: Feminine, withZeros: false}, want: "два квадриллиона две"},
		{name: "2000000000000002ff", args: args{number: 2000000000000002, gender: Neuter, withZeros: true},
			want: "два квадриллиона ноль триллионов ноль миллиардов ноль миллионов ноль тысяч два"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberInWords(tt.args.number, tt.args.gender, tt.args.withZeros); got != tt.want {
				t.Errorf("NumberInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
