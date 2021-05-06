package pathhelper

import (
	"reflect"
	"testing"
)

func TestSplitPath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "00", args: args{p: ""}, want: []string{}},
		{name: "01", args: args{p: "/"}, want: []string{}},
		{name: "02", args: args{p: "//"}, want: []string{}},
		{name: "1", args: args{p: "a/b/c.d"}, want: []string{"a", "b", "c.d"}},
		{name: "2", args: args{p: "/a/"}, want: []string{"a"}},
		{name: "3", args: args{p: "a/"}, want: []string{"a"}},
		{name: "4", args: args{p: "/a"}, want: []string{"a"}},
		{name: "5", args: args{p: "a"}, want: []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPath(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartSlash(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{p: ""}, want: "/"},
		{name: "1", args: args{p: "a/b/c.d"}, want: "/a/b/c.d"},
		{name: "2", args: args{p: "/a/b/c.d"}, want: "/a/b/c.d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartSlash(tt.args.p); got != tt.want {
				t.Errorf("StartSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndSlash(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{p: ""}, want: "/"},
		{name: "1", args: args{p: "/a/b/c.d"}, want: "/a/b/c.d/"},
		{name: "2", args: args{p: "a/b/c.d/"}, want: "a/b/c.d/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndSlash(tt.args.p); got != tt.want {
				t.Errorf("EndSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
