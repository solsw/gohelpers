package russian

import "testing"

func TestFloatInWords(t *testing.T) {
	type args struct {
		f            float64
		fraction     Fraction
		binder       Binder
		showZeroInt  bool
		showZeroFrac bool
		withZeroInt  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "01", args: args{f: 0.0},
			want: "ноль"},
		{name: "02",
			args: args{f: 0.0,
				fraction:     Tenth,
				showZeroFrac: true},
			want: "ноль десятых"},
		{name: "03",
			args: args{
				f:            0.0,
				fraction:     Tenth,
				showZeroInt:  true,
				showZeroFrac: true},
			want: "ноль ноль десятых"},
		{name: "04",
			args: args{
				f:            -0.0,
				fraction:     Tenth,
				binder:       And,
				showZeroInt:  true,
				showZeroFrac: true},
			want: "ноль и ноль десятых"},
		{name: "05",
			args: args{
				f:            0.0,
				fraction:     Tenth,
				binder:       Whole,
				showZeroInt:  true,
				showZeroFrac: true},
			want: "ноль целых ноль десятых"},
		{name: "06",
			args: args{
				f:            0.0,
				fraction:     Thousandth,
				showZeroFrac: true},
			want: "ноль тысячных"},
		{name: "0.1_10",
			args: args{
				f:            0.1,
				fraction:     Tenth,
				showZeroFrac: true},
			want: "одна десятая"},
		{name: "-0.1_100",
			args: args{
				f:            -0.1,
				fraction:     Hundredth,
				showZeroFrac: true},
			want: "минус десять сотых"},
		{name: "0.1_1000",
			args: args{
				f:            0.1,
				fraction:     Thousandth,
				showZeroFrac: true},
			want: "сто тысячных"},
		{name: "21.123_100_Whole",
			args: args{
				f:            21.123,
				fraction:     Hundredth,
				binder:       Whole,
				showZeroFrac: true},
			want: "двадцать одна целая двенадцать сотых"},
		{name: "-21.123_1000000000_And",
			args: args{
				f:            -21.123,
				fraction:     Milliardth,
				binder:       And,
				showZeroFrac: true},
			want: "минус двадцать один и сто двадцать три миллиона миллиардных"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatInWords(tt.args.f, tt.args.fraction, tt.args.binder,
				tt.args.showZeroInt, tt.args.showZeroFrac, tt.args.withZeroInt); got != tt.want {
				t.Errorf("FloatInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
