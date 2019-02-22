package filepathhelper

import (
	"reflect"
	"testing"
)

func TestSplitFilePath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{p: "a\\b\\c.d"}, want: []string{"a", "b", "c.d"}},
		{name: "2", args: args{p: "a\\b\\"}, want: []string{"a", "b"}},
		{name: "3", args: args{p: "a\\b"}, want: []string{"a", "b"}},
		{name: "4", args: args{p: "a\\"}, want: []string{"a"}},
		{name: "5", args: args{p: "a"}, want: []string{"a"}},
		{name: "6", args: args{p: "\\"}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitFilePath(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseNoExt(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{p: "a\\b\\c.d"}, want: "c"},
		{name: "2", args: args{p: "a\\b\\cd.e"}, want: "cd"},
		{name: "3", args: args{p: "a\\b\\c."}, want: "c"},
		{name: "4", args: args{p: "a\\b\\."}, want: ""},
		{name: "4", args: args{p: "a\\b\\"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseNoExt(tt.args.p); got != tt.want {
				t.Errorf("BaseNoExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartSeparator(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{p: ""}, want: ""},
		{name: "1", args: args{p: "a\\b\\c.d"}, want: "\\a\\b\\c.d"},
		{name: "2", args: args{p: "\\a\\b\\c.d"}, want: "\\a\\b\\c.d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartSeparator(tt.args.p); got != tt.want {
				t.Errorf("StartSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndSeparator(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{p: ""}, want: ""},
		{name: "1", args: args{p: "\\a\\b\\c.d"}, want: "\\a\\b\\c.d\\"},
		{name: "2", args: args{p: "a\\b\\c.d\\"}, want: "a\\b\\c.d\\"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndSeparator(tt.args.p); got != tt.want {
				t.Errorf("EndSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
