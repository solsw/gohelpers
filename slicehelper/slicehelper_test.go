package slicehelper

import (
	"reflect"
	"testing"
)

func TestConcatElSl(t *testing.T) {
	type args struct {
		el interface{}
		sl []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{name: "01", args: args{el: nil, sl: nil}, want: []interface{}{nil}},
		{name: "02", args: args{el: nil, sl: []interface{}{}}, want: []interface{}{nil}},
		{name: "1", args: args{el: 1, sl: nil}, want: []interface{}{1}},
		{name: "2", args: args{el: nil, sl: []interface{}{"qwerty"}}, want: []interface{}{nil, "qwerty"}},
		{name: "3", args: args{el: 1, sl: []interface{}{"qwerty"}}, want: []interface{}{1, "qwerty"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatElSl(tt.args.el, tt.args.sl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcatElSl() = %v, want %v", got, tt.want)
			}
		})
	}
}

var eqInt = func(v1, v2 interface{}) bool {
	return v1.(int) == v2.(int)
}

func TestContainsEqEr(t *testing.T) {
	type args struct {
		sl []interface{}
		el interface{}
		eq func(interface{}, interface{}) bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "e1", args: args{sl: nil, el: nil, eq: nil}, wantErr: true},
		{name: "e2", args: args{sl: []interface{}{1}, el: 1, eq: nil}, wantErr: true},
		{name: "01", args: args{sl: nil, el: nil, eq: eqInt}, want: false},
		{name: "02", args: args{sl: nil, el: 1, eq: eqInt}, want: false},
		{name: "1", args: args{sl: []interface{}{1}, el: 1, eq: eqInt}, want: true},
		{name: "2", args: args{sl: []interface{}{1, 2, 3}, el: 1, eq: eqInt}, want: true},
		{name: "3", args: args{sl: []interface{}{1, 2, 3}, el: 2, eq: eqInt}, want: true},
		{name: "4", args: args{sl: []interface{}{1, 2, 3}, el: 3, eq: eqInt}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ContainsEqEr(tt.args.sl, tt.args.el, tt.args.eq)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContainsEqEr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContainsEqEr() = %v, want %v", got, tt.want)
			}
		})
	}
}

var cmpInt = func(v1, v2 interface{}) int {
	i1 := v1.(int)
	i2 := v2.(int)
	if i1 < i2 {
		return -1
	}
	if i1 > i2 {
		return 1
	}
	return 0
}

func TestContainsCmp(t *testing.T) {
	type args struct {
		sl  []interface{}
		el  interface{}
		cmp func(interface{}, interface{}) int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "01", args: args{sl: nil, el: nil, cmp: cmpInt}, want: false},
		{name: "02", args: args{sl: nil, el: 1, cmp: cmpInt}, want: false},
		{name: "1", args: args{sl: []interface{}{1}, el: 1, cmp: cmpInt}, want: true},
		{name: "2", args: args{sl: []interface{}{1, 2, 3}, el: 1, cmp: cmpInt}, want: true},
		{name: "3", args: args{sl: []interface{}{1, 2, 3}, el: 2, cmp: cmpInt}, want: true},
		{name: "4", args: args{sl: []interface{}{1, 2, 3}, el: 3, cmp: cmpInt}, want: true},
		{name: "5", args: args{sl: []interface{}{1, 3, 5}, el: 0, cmp: cmpInt}, want: false},
		{name: "6", args: args{sl: []interface{}{1, 3, 5}, el: 2, cmp: cmpInt}, want: false},
		{name: "7", args: args{sl: []interface{}{1, 3, 5}, el: 4, cmp: cmpInt}, want: false},
		{name: "8", args: args{sl: []interface{}{1, 3, 5}, el: 6, cmp: cmpInt}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsCmp(tt.args.sl, tt.args.el, tt.args.cmp); got != tt.want {
				t.Errorf("ContainsCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAt(t *testing.T) {
	var sl0 []interface{}
	type args struct {
		sl  []interface{}
		idx int
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{name: "1e", args: args{sl: sl0}, want: nil, wantErr: true},
		{name: "2e", args: args{sl: []interface{}{}}, want: nil, wantErr: true},
		{name: "3e", args: args{sl: []interface{}{1, 2}, idx: 2}, wantErr: true},
		{name: "1i", args: args{sl: []interface{}{1}}, want: []interface{}{}},
		{name: "1s", args: args{sl: []interface{}{"one"}}, want: []interface{}{}},
		{name: "20i", args: args{sl: []interface{}{1, 2}}, want: []interface{}{2}},
		{name: "21i", args: args{sl: []interface{}{1, 2}, idx: 1}, want: []interface{}{1}},
		{name: "31i", args: args{sl: []interface{}{1, 2, 3}, idx: 1}, want: []interface{}{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveAt(tt.args.sl, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAtInPlace(t *testing.T) {
	type args struct {
		sl  *[]interface{}
		idx int
	}
	tests := []struct {
		name    string
		args    args
		want    *[]interface{}
		wantErr bool
	}{
		{name: "0", args: args{sl: &[]interface{}{1, 2, 3}, idx: 0}, want: &[]interface{}{2, 3}},
		{name: "1", args: args{sl: &[]interface{}{1, 2, 3}, idx: 1}, want: &[]interface{}{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveAtInPlace(tt.args.sl, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAtInPlace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAtInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
