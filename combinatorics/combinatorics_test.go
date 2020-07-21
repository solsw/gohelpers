package combinatorics

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	var nilSl []interface{}
	type args struct {
		sl []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    [][]interface{}
		wantErr bool
	}{
		{name: "1e", args: args{sl: nilSl}, want: nil, wantErr: true},
		{name: "2e", args: args{sl: []interface{}{}}, want: nil, wantErr: true},
		{name: "1i",
			args: args{sl: []interface{}{1}},
			want: [][]interface{}{{1}}},
		{name: "1s",
			args: args{sl: []interface{}{"one"}},
			want: [][]interface{}{{"one"}}},
		{name: "2i",
			args: args{sl: []interface{}{1, 2}},
			want: [][]interface{}{{1, 2}, {2, 1}}},
		{name: "2m",
			args: args{sl: []interface{}{"one", 2}},
			want: [][]interface{}{{"one", 2}, {2, "one"}}},
		{name: "3i",
			args: args{sl: []interface{}{1, 2, 3}},
			want: [][]interface{}{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Permutations(tt.args.sl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Permutations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
