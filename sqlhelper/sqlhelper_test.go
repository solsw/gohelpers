package sqlhelper

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestStrToNullStr(t *testing.T) {
	type args struct {
		s           string
		emptyIsNULL bool
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{name: "1", args: args{s: "qwerty"}, want: sql.NullString{String: "qwerty", Valid: true}},
		{name: "2", args: args{s: ""}, want: sql.NullString{String: "", Valid: true}},
		{name: "3", args: args{s: "", emptyIsNULL: true}, want: sql.NullString{String: "", Valid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToNullStr(tt.args.s, tt.args.emptyIsNULL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToNullStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullStrToStr(t *testing.T) {
	type args struct {
		ns sql.NullString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{ns: sql.NullString{String: "qwerty", Valid: true}}, want: "qwerty"},
		{name: "2", args: args{ns: sql.NullString{String: "", Valid: true}}, want: ""},
		{name: "3", args: args{ns: sql.NullString{String: "qwerty", Valid: false}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullStrToStr(tt.args.ns); got != tt.want {
				t.Errorf("NullStrToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
