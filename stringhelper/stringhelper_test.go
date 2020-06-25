package stringhelper

import (
	"math/rand"
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/solsw/gohelpers/oshelper"
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

func TestNthRuneStrict(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{name: "1e", args: args{s: "", n: 2}, want: utf8.RuneError, wantErr: true},
		{name: "invalidString", args: args{s: invalidString}, want: utf8.RuneError, wantErr: true},
		{name: "2e", args: args{s: " ", n: 2}, want: utf8.RuneError, wantErr: true},
		{name: "1", args: args{s: "3.1415926535", n: 2}, want: '1', wantErr: false},
		{name: "2", args: args{s: "йцукен", n: 2}, want: 'у', wantErr: false},
		{name: "3", args: args{s: invalidString, n: 2}, want: utf8.RuneError, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NthRuneStrict(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NthRuneStrict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NthRuneStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNthRuneAny(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{name: "1e", args: args{s: "", n: 2}, want: utf8.RuneError, wantErr: true},
		{name: "2e", args: args{s: " ", n: 2}, want: utf8.RuneError, wantErr: true},
		{name: "invalidString1", args: args{s: invalidString}, want: 'ф', wantErr: false},
		{name: "invalidString2", args: args{s: invalidString, n: 2}, want: 'в', wantErr: false},
		{name: "invalidString3", args: args{s: invalidString, n: 8}, want: utf8.RuneError, wantErr: true},
		{name: "1", args: args{s: "3.1415926535", n: 2}, want: '1', wantErr: false},
		{name: "2", args: args{s: "йцукен", n: 2}, want: 'у', wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NthRuneAny(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NthRuneAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NthRuneAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNthWord(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, want: "", wantErr: true},
		{name: "2e", args: args{s: "  "}, want: "", wantErr: true},
		{name: "3e", args: args{s: "q w", n: 2}, want: "", wantErr: true},
		{name: "1", args: args{s: "q w", n: 0}, want: "q", wantErr: false},
		{name: "2", args: args{s: "q w", n: 1}, want: "w", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NthWord(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NthWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NthWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, want: "", wantErr: true},
		{name: "2e", args: args{s: "  "}, want: "", wantErr: true},
		{name: "1", args: args{s: "q w"}, want: "w", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastWord(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNthWordDelims(t *testing.T) {
	type args struct {
		s      string
		n      uint
		delims []rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, want: "", wantErr: true},
		{name: "2e", args: args{s: "  "}, want: "", wantErr: true},
		{name: "3e", args: args{s: "q w", n: 2, delims: []rune{}}, want: "", wantErr: true},
		{name: "1", args: args{s: "q w", n: 1}, want: "w", wantErr: false},
		{name: "2", args: args{s: "q-w", n: 1, delims: []rune{'-', '+', '_'}}, want: "w", wantErr: false},
		{name: "3", args: args{s: "q-w+-e", n: 2, delims: []rune{'-', '+', '_'}}, want: "e", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NthWordDelims(tt.args.s, tt.args.n, tt.args.delims)
			if (err != nil) != tt.wantErr {
				t.Errorf("NthWordDelims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NthWordDelims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastWordDelims(t *testing.T) {
	type args struct {
		s      string
		delims []rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, wantErr: true},
		{name: "2e", args: args{s: "  "}, wantErr: true},
		{name: "1", args: args{s: "q w"}, want: "w", wantErr: false},
		{name: "2", args: args{s: "q-w", delims: []rune{'-', '+', '_'}}, want: "w", wantErr: false},
		{name: "3", args: args{s: "q-w+-e", delims: []rune{'-', '+', '_'}}, want: "e", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastWordDelims(tt.args.s, tt.args.delims)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastWordDelims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastWordDelims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstr(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{start: -1}, wantErr: true},
		{name: "2e", args: args{length: -2}, wantErr: true},
		{name: "3e", args: args{s: "1", start: 0, length: 2}, wantErr: true},
		{name: "4e", args: args{s: "1", start: 1, length: 1}, wantErr: true},
		{name: "5e", args: args{s: "1", start: 2, length: 0}, wantErr: true},
		{name: "1", args: args{s: "1", start: 0, length: 1}, want: "1"},
		{name: "2", args: args{s: "1234", start: 0, length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", start: 0, length: 2}, want: "12"},
		{name: "4", args: args{s: "1234", start: 1, length: 3}, want: "234"},
		{name: "5", args: args{s: "1234", start: 1, length: 0}, want: ""},
		{name: "6", args: args{s: "1234", start: 4, length: 0}, want: ""},
		{name: "1r", args: args{s: "йцукен", start: 1, length: 3}, want: "цук"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Substr(tt.args.s, tt.args.start, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Substring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Substring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrBeg(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{length: -1}, wantErr: true},
		{name: "2e", args: args{s: "1", length: 2}, wantErr: true},
		{name: "0", args: args{s: "", length: 0}, want: ""},
		{name: "1", args: args{s: "1", length: 0}, want: ""},
		{name: "2", args: args{s: "1", length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", length: 0}, want: ""},
		{name: "4", args: args{s: "1234", length: 1}, want: "1"},
		{name: "5", args: args{s: "1234", length: 2}, want: "12"},
		{name: "6", args: args{s: "1234", length: 3}, want: "123"},
		{name: "7", args: args{s: "1234", length: 4}, want: "1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrBeg(tt.args.s, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrBeg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrBeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrEnd(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{length: -1}, wantErr: true},
		{name: "2e", args: args{s: "1", length: 2}, wantErr: true},
		{name: "0", args: args{s: "", length: 0}, want: ""},
		{name: "1", args: args{s: "1", length: 0}, want: ""},
		{name: "2", args: args{s: "1", length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", length: 0}, want: ""},
		{name: "4", args: args{s: "1234", length: 1}, want: "4"},
		{name: "5", args: args{s: "1234", length: 2}, want: "34"},
		{name: "6", args: args{s: "1234", length: 3}, want: "234"},
		{name: "7", args: args{s: "1234", length: 4}, want: "1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrEnd(tt.args.s, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrToEnd(t *testing.T) {
	type args struct {
		s     string
		start int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e", args: args{start: -1}, wantErr: true},
		{name: "2e", args: args{s: "1", start: 2}, wantErr: true},
		{name: "1", args: args{s: "1", start: 1}, want: ""},
		{name: "2", args: args{s: "1", start: 0}, want: "1"},
		{name: "3", args: args{s: "1234", start: 0}, want: "1234"},
		{name: "4", args: args{s: "1234", start: 2}, want: "34"},
		{name: "5", args: args{s: "1234", start: 3}, want: "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrToEnd(tt.args.s, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrToEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastByte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, wantErr: true},
		{name: "1", args: args{s: "qwerty"}, want: 'y'},
		{name: "2", args: args{s: "йцукен"}, want: 189},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastByte(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastRune(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{name: "1e", args: args{s: ""}, want: utf8.RuneError, wantErr: true},
		{name: "1", args: args{s: "qwerty"}, want: rune('y')},
		{name: "2", args: args{s: "йцукен"}, want: rune('н')},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastRune(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1t", args: args{ss: nil}, want: nil},
		{name: "2t", args: args{ss: []string{}}, want: []string{}},
		{name: "3t", args: args{ss: []string{"qwerty"}}, want: []string{"qwerty"}},
		{name: "1", args: args{ss: []string{"one", "two", "three", "four"}}, want: []string{"one", "two", "three", "four"}},
		{name: "2", args: args{ss: []string{"one", "two", "three", "one", "two"}}, want: []string{"one", "two", "three"}},
		{name: "3",
			args: args{ss: []string{"one", "two", "three", "one", "two", "two", "three"}},
			want: []string{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueSorted(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1t", args: args{ss: nil}, want: nil},
		{name: "2t", args: args{ss: []string{}}, want: []string{}},
		{name: "3t", args: args{ss: []string{"qwerty"}}, want: []string{"qwerty"}},
		{name: "1", args: args{ss: []string{"one", "two", "three", "four"}}, want: []string{"four", "one", "three", "two"}},
		{name: "2", args: args{ss: []string{"one", "two", "three", "one", "two"}}, want: []string{"one", "three", "two"}},
		{name: "3",
			args: args{ss: []string{"one", "two", "three", "one", "two", "two", "three"}},
			want: []string{"one", "three", "two"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueSorted(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func initSS() []string {
	s := "qwertyuioplkjhgfdsazxcvbnm"
	rr := []rune(s)
	var ss []string
	rnd := rand.New(rand.NewSource(1234))
	for i := 0; i < 1000; i++ {
		rnd.Shuffle(len(rr), func(i, j int) { rr[i], rr[j] = rr[j], rr[i] })
		ss = append(ss, string(rr))
	}
	ss = append(ss, ss...)
	return ss
}

func BenchmarkUnique(b *testing.B) {
	ss := initSS()
	for i := 0; i < b.N; i++ {
		Unique(ss)
	}
}

func BenchmarkUniqueSorted(b *testing.B) {
	ss := initSS()
	for i := 0; i < b.N; i++ {
		UniqueSorted(ss)
	}
}

func TestRemoveEscSGR(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{s: ""}, want: ""},
		{name: "1", args: args{s: "23"}, want: "23"},
		{name: "2", args: args{s: "\x1B[36m"}, want: ""},
		{name: "3", args: args{s: "\x1B[36mINFO"}, want: "INFO"},
		{name: "4", args: args{s: "\x1B[36mINFO\x1B[0m"}, want: "INFO"},
		{name: "5", args: args{s: "\x1B[36mINFO\x1B[0m[0000]"}, want: "INFO[0000]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEscSGR(tt.args.s); got != tt.want {
				t.Errorf("RemoveEscSGR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjustNewLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{s: ""}, want: ""},
		{name: "01", args: args{s: "\n"}, want: oshelper.NewLine},
		{name: "02", args: args{s: "\r"}, want: oshelper.NewLine},
		{name: "03", args: args{s: "\r\n"}, want: oshelper.NewLine},
		{name: "04", args: args{s: "\n\r"}, want: oshelper.NewLine + oshelper.NewLine},
		{name: "1", args: args{s: "23"}, want: "23"},
		{name: "2", args: args{s: "2\r\n3"}, want: "2" + oshelper.NewLine + "3"},
		{name: "3", args: args{s: "2\r3\r4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "4", args: args{s: "2\n3\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "5", args: args{s: "2\r3\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "6", args: args{s: "2\n3\r4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "7", args: args{s: "2\n3\n"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "8", args: args{s: "2\r3\r"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "9", args: args{s: "2\n3\r"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "9", args: args{s: "2\r3\n"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "10", args: args{s: "2\r3\r\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjustNewLines(tt.args.s); got != tt.want {
				t.Errorf("AdjustNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToStrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "0", args: args{s: ""}, want: []string{""}},
		{name: "1", args: args{s: "2\r3\r\n4"}, want: []string{"2", "3", "4"}},
		{name: "3", args: args{s: "2\r3\r\n4\r"}, want: []string{"2", "3", "4", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveLastStringIfEmpty(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "0", args: args{ss: nil}, want: nil},
		{name: "1", args: args{ss: []string{}}, want: []string{}},
		{name: "2", args: args{ss: []string{"1", ""}}, want: []string{"1"}},
		{name: "3", args: args{ss: []string{"1", "", ""}}, want: []string{"1", ""}},
		{name: "4", args: args{ss: []string{"1", "", "2"}}, want: []string{"1", "", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLastStringIfEmpty(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLastStringIfEmpty() = %v, want %v", got, tt.want)
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
