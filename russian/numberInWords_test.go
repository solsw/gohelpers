package russian

import "testing"

func TestNumberInWords(t *testing.T) {
	type args struct {
		number   int64
		gender   GrammaticalGender
		withZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0mt", args: args{number: 0, gender: Masculine, withZero: true},
			want: "ноль"},
		{name: "0mf", args: args{number: 0, gender: Masculine, withZero: false},
			want: "ноль"},
		{name: "-1mt", args: args{number: -1, gender: Masculine, withZero: true},
			want: "минус один"},
		{name: "1ft", args: args{number: 1, gender: Feminine, withZero: true},
			want: "одна"},
		{name: "100ff", args: args{number: 100, gender: Feminine, withZero: false},
			want: "сто"},
		{name: "1000mt", args: args{number: 1000, gender: Masculine, withZero: true},
			want: "одна тысяча ноль"},
		{name: "1000mf", args: args{number: 1000, gender: Masculine, withZero: false},
			want: "одна тысяча"},
		{name: "1001mf", args: args{number: 1001, gender: Masculine, withZero: false},
			want: "одна тысяча один"},
		{name: "2000002ft", args: args{number: 2000002, gender: Feminine, withZero: true},
			want: "два миллиона ноль тысяч две"},
		{name: "2000002ff", args: args{number: 2000002, gender: Feminine, withZero: false},
			want: "два миллиона две"},
		{name: "2000000000000002ff", args: args{number: 2000000000000002, gender: Feminine, withZero: false},
			want: "два квадриллиона две"},
		{name: "2000000000000002nt", args: args{number: 2000000000000002, gender: Neuter, withZero: true},
			want: "два квадриллиона ноль триллионов ноль миллиардов ноль миллионов ноль тысяч два"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberInWords(tt.args.number, tt.args.withZero, tt.args.gender); got != tt.want {
				t.Errorf("NumberInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
