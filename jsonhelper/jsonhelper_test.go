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

func TestMustFormatStrToStrDef(t *testing.T) {
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
			if got := MustFormatStrToStrDef(tt.args.json); got != tt.want {
				t.Errorf("MustFormatStrToStrDef() = %v, want %v", got, tt.want)
			}
		})
	}
}
