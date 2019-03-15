package filepathhelper

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestNoExt(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{p: ``}, want: ``},
		{name: "01", args: args{p: `.`}, want: ``},
		{name: "02", args: args{p: `.a`}, want: ``},
		{name: "1", args: args{p: `a/b/c`}, want: `a/b/c`},
		{name: "2", args: args{p: `a/b/c.d`}, want: `a/b/c`},
		{name: "3", args: args{p: `a/b/cd.e`}, want: `a/b/cd`},
		{name: "4", args: args{p: `a/b/c.`}, want: `a/b/c`},
		{name: "5", args: args{p: `a/b/.c`}, want: `a/b/`},
		{name: "6", args: args{p: `a/b/.`}, want: `a/b/`},
		{name: "7", args: args{p: `a/b/`}, want: `a/b/`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoExt(filepath.FromSlash(tt.args.p)); got != filepath.FromSlash(tt.want) {
				t.Errorf("NoExt() = %v, want %v", got, tt.want)
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
		{name: "00", args: args{p: ``}, want: ``},
		{name: "01", args: args{p: `.`}, want: ``},
		{name: "02", args: args{p: `.a`}, want: ``},
		{name: "1", args: args{p: `a/b/c`}, want: `c`},
		{name: "2", args: args{p: `a/b/c.d`}, want: `c`},
		{name: "3", args: args{p: `a/b/cd.e`}, want: `cd`},
		{name: "4", args: args{p: `a/b/c.`}, want: `c`},
		{name: "5", args: args{p: `a/b/.c`}, want: ``},
		{name: "6", args: args{p: `a/b/.`}, want: ``},
		{name: "7", args: args{p: `a/b/`}, want: ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseNoExt(filepath.FromSlash(tt.args.p)); got != tt.want {
				t.Errorf("BaseNoExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeExt(t *testing.T) {
	type args struct {
		p   string
		ext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{p: ``, ext: ``}, want: ``},
		{name: "01", args: args{p: ``, ext: `txt`}, want: ``},
		{name: "1", args: args{p: `a/b/c`, ext: ``}, want: `a/b/c`},
		{name: "2", args: args{p: `a/b/c`, ext: `.`}, want: `a/b/c`},
		{name: "3", args: args{p: `.`, ext: `txt`}, want: `.txt`},
		{name: "4", args: args{p: `.`, ext: `.txt`}, want: `.txt`},
		{name: "5", args: args{p: `a/b/c`, ext: `d`}, want: `a/b/c.d`},
		{name: "6", args: args{p: `a/b/c`, ext: `.d`}, want: `a/b/c.d`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeExt(filepath.FromSlash(tt.args.p), tt.args.ext); got != filepath.FromSlash(tt.want) {
				t.Errorf("ChangeExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitFilePath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{p: "a/b/c.d"}, want: []string{"a", "b", "c.d"}},
		{name: "2", args: args{p: "a/b/"}, want: []string{"a", "b"}},
		{name: "3", args: args{p: "a/b"}, want: []string{"a", "b"}},
		{name: "4", args: args{p: "a/"}, want: []string{"a"}},
		{name: "5", args: args{p: "a"}, want: []string{"a"}},
		{name: "6", args: args{p: "/"}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitFilePath(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitFilePath() = %v, want %v", got, tt.want)
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
		{name: "0", args: args{p: ""}, want: "/"},
		{name: "1", args: args{p: "a/b/c.d"}, want: "/a/b/c.d"},
		{name: "2", args: args{p: "/a/b/c.d"}, want: "/a/b/c.d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartSeparator(filepath.FromSlash(tt.args.p)); got != filepath.FromSlash(tt.want) {
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
		{name: "0", args: args{p: ""}, want: "/"},
		{name: "1", args: args{p: "/a/b/c.d"}, want: "/a/b/c.d/"},
		{name: "2", args: args{p: "a/b/c.d/"}, want: "a/b/c.d/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndSeparator(filepath.FromSlash(tt.args.p)); got != filepath.FromSlash(tt.want) {
				t.Errorf("EndSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
