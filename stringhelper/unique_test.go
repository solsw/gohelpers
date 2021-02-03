package stringhelper

import (
	"math/rand"
	"reflect"
	"testing"
)

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
