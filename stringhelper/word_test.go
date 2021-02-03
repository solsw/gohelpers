package stringhelper

import (
	"testing"
)

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

func TestNthWordFunc(t *testing.T) {
	type args struct {
		s string
		n uint
		f func(rune) bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "0", args: args{s: "q0w00e000", n: 1}, wantErr: true},
		{name: "1", args: args{s: "q0w00e000", n: 1, f: func(r rune) bool { return r == '0' }}, want: "w"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NthWordFunc(tt.args.s, tt.args.n, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("NthWordFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NthWordFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastWordFunc(t *testing.T) {
	type args struct {
		s string
		f func(rune) bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "0", args: args{s: "q0w00e000"}, wantErr: true},
		{name: "1", args: args{s: "q0w00e000", f: func(r rune) bool { return r == '0' }}, want: "e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastWordFunc(tt.args.s, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastWordFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastWordFunc() = %v, want %v", got, tt.want)
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
