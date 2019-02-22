package slicehelper

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	var sl0 []interface{}
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
		{name: "1e", args: args{sl: &sl0}, want: nil, wantErr: true},
		{name: "2e", args: args{sl: &[]interface{}{}}, want: nil, wantErr: true},
		{name: "3e", args: args{sl: &[]interface{}{1, 2}, idx: 2}, wantErr: true},
		{name: "1i", args: args{sl: &[]interface{}{1}}, want: &[]interface{}{}},
		{name: "1s", args: args{sl: &[]interface{}{"one"}}, want: &[]interface{}{}},
		{name: "20i", args: args{sl: &[]interface{}{1, 2}}, want: &[]interface{}{2}},
		{name: "21i", args: args{sl: &[]interface{}{1, 2}, idx: 1}, want: &[]interface{}{1}},
		{name: "31i", args: args{sl: &[]interface{}{1, 2, 3}, idx: 1}, want: &[]interface{}{1, 3}},
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
