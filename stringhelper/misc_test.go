package stringhelper

import (
	"reflect"
	"testing"

	"github.com/solsw/gohelpers/oshelper"
)

func TestAdjustNewLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00", args: args{s: ""}, want: ""},
		{name: "01", args: args{s: "\n"}, want: oshelper.NewLine},
		{name: "02", args: args{s: "\r"}, want: oshelper.NewLine},
		{name: "03", args: args{s: "\r\n"}, want: oshelper.NewLine},
		{name: "04", args: args{s: "\n\r"}, want: oshelper.NewLine + oshelper.NewLine},
		{name: "1", args: args{s: "23"}, want: "23"},
		{name: "2", args: args{s: "2\r\n3"}, want: "2" + oshelper.NewLine + "3"},
		{name: "3", args: args{s: "2\r3\r4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "4", args: args{s: "2\n3\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "5", args: args{s: "2\r3\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "6", args: args{s: "2\n3\r4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
		{name: "7", args: args{s: "2\n3\n"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "8", args: args{s: "2\r3\r"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "9", args: args{s: "2\n3\r"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "9", args: args{s: "2\r3\n"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine},
		{name: "10", args: args{s: "2\r3\r\n4"}, want: "2" + oshelper.NewLine + "3" + oshelper.NewLine + "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjustNewLines(tt.args.s); got != tt.want {
				t.Errorf("AdjustNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToStrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "0", args: args{s: ""}, want: []string{""}},
		{name: "1", args: args{s: "2\r3\r\n4"}, want: []string{"2", "3", "4"}},
		{name: "3", args: args{s: "2\r3\r\n4\r"}, want: []string{"2", "3", "4", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveLastStringIfEmpty(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "0", args: args{ss: nil}, want: nil},
		{name: "1", args: args{ss: []string{}}, want: []string{}},
		{name: "2", args: args{ss: []string{"1", ""}}, want: []string{"1"}},
		{name: "3", args: args{ss: []string{"1", "", ""}}, want: []string{"1", ""}},
		{name: "4", args: args{ss: []string{"1", "", "2"}}, want: []string{"1", "", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLastStringIfEmpty(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLastStringIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveSGREsc(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{s: ""}, want: ""},
		{name: "1", args: args{s: "23"}, want: "23"},
		{name: "2", args: args{s: "\x1B[36m"}, want: ""},
		{name: "3", args: args{s: "\x1B[36mINFO"}, want: "INFO"},
		{name: "4", args: args{s: "\x1B[36mINFO\x1B[0m"}, want: "INFO"},
		{name: "5", args: args{s: "\x1B[36mINFO\x1B[0m[0000]"}, want: "INFO[0000]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSGREsc(tt.args.s); got != tt.want {
				t.Errorf("RemoveSGREsc() = %v, want %v", got, tt.want)
			}
		})
	}
}
