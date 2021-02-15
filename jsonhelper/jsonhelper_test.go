package jsonhelper

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestFormatDef(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{name: "1e", args: args{r: &strings.Reader{}}, wantErr: true},
		{name: "1", args: args{r: strings.NewReader(`{"i":1,"s":"string"}`)},
			wantW: "{\n  \"i\": 1,\n  \"s\": \"string\"\n}\n"},
		{name: "2", args: args{r: strings.NewReader(`[{"i":1,"s":"one"}]`)},
			wantW: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  }\n]\n"},
		{name: "3", args: args{r: strings.NewReader(`[{"i":1,"s":"one"},{"i":2,"s":"two"}]`)},
			wantW: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  },\n  {\n    \"i\": 2,\n    \"s\": \"two\"\n  }\n]\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := FormatDef(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("FormatDef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("FormatDef() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestFormatStrToStrDefMust(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{json: `{"i":1,"s":"string"}`}, want: "{\n  \"i\": 1,\n  \"s\": \"string\"\n}\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatStrToStrDefMust(tt.args.json); got != tt.want {
				t.Errorf("FormatStrToStrDefMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatToStrDefMust(t *testing.T) {
	type is struct {
		I int
		S string
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{v: is{1, "one"}},
			want: `{
  "I": 1,
  "S": "one"
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatToStrDefMust(tt.args.v); got != tt.want {
				t.Errorf("FormatToStrDefMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
