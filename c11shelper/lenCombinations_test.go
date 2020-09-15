package c11shelper

import (
	"reflect"
	"testing"
)

func TestLenCombinations(t *testing.T) {
	type args struct {
		arrLen  int
		combLen int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "20", args: args{arrLen: 2, combLen: 0}, want: nil},
		{name: "12", args: args{arrLen: 1, combLen: 2}, want: nil},
		{name: "11", args: args{arrLen: 1, combLen: 1}, want: [][]int{{0}}},
		{name: "21", args: args{arrLen: 2, combLen: 1}, want: [][]int{{0}, {1}}},
		{name: "41", args: args{arrLen: 4, combLen: 1}, want: [][]int{{0}, {1}, {2}, {3}}},
		{name: "22", args: args{arrLen: 2, combLen: 2}, want: [][]int{{0, 1}}},
		{name: "32", args: args{arrLen: 3, combLen: 2}, want: [][]int{{0, 1}, {0, 2}, {1, 2}}},
		{name: "42", args: args{arrLen: 4, combLen: 2}, want: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}},
		{name: "43", args: args{arrLen: 4, combLen: 3}, want: [][]int{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3}}},
		{name: "44", args: args{arrLen: 4, combLen: 4}, want: [][]int{{0, 1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenCombinations(tt.args.arrLen, tt.args.combLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LenCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
