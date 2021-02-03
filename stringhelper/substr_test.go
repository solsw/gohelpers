package stringhelper

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "3e", args: args{s: "1", start: 0, length: 2}, wantErr: true},
		{name: "4e", args: args{s: "1", start: 1, length: 1}, wantErr: true},
		{name: "5e", args: args{s: "1", start: 2, length: 0}, wantErr: true},
		{name: "1", args: args{s: "1", start: 0, length: 1}, want: "1"},
		{name: "2", args: args{s: "1234", start: 0, length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", start: 0, length: 2}, want: "12"},
		{name: "4", args: args{s: "1234", start: 1, length: 3}, want: "234"},
		{name: "5", args: args{s: "1234", start: 1, length: 0}, want: ""},
		{name: "6", args: args{s: "1234", start: 4, length: 0}, want: ""},
		{name: "1r", args: args{s: "йцукен", start: 1, length: 3}, want: "цук"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Substr(tt.args.s, tt.args.start, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Substring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Substring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrBeg(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "2e", args: args{s: "1", length: 2}, wantErr: true},
		{name: "0", args: args{s: "", length: 0}, want: ""},
		{name: "1", args: args{s: "1", length: 0}, want: ""},
		{name: "2", args: args{s: "1", length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", length: 0}, want: ""},
		{name: "4", args: args{s: "1234", length: 1}, want: "1"},
		{name: "5", args: args{s: "1234", length: 2}, want: "12"},
		{name: "6", args: args{s: "1234", length: 3}, want: "123"},
		{name: "7", args: args{s: "1234", length: 4}, want: "1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrBeg(tt.args.s, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrBeg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrBeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrEnd(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "2e", args: args{s: "1", length: 2}, wantErr: true},
		{name: "0", args: args{s: "", length: 0}, want: ""},
		{name: "1", args: args{s: "1", length: 0}, want: ""},
		{name: "2", args: args{s: "1", length: 1}, want: "1"},
		{name: "3", args: args{s: "1234", length: 0}, want: ""},
		{name: "4", args: args{s: "1234", length: 1}, want: "4"},
		{name: "5", args: args{s: "1234", length: 2}, want: "34"},
		{name: "6", args: args{s: "1234", length: 3}, want: "234"},
		{name: "7", args: args{s: "1234", length: 4}, want: "1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrEnd(tt.args.s, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrToEnd(t *testing.T) {
	type args struct {
		s     string
		start int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "2e", args: args{s: "1", start: 2}, wantErr: true},
		{name: "1", args: args{s: "1", start: 1}, want: ""},
		{name: "2", args: args{s: "1", start: 0}, want: "1"},
		{name: "3", args: args{s: "1234", start: 0}, want: "1234"},
		{name: "4", args: args{s: "1234", start: 2}, want: "34"},
		{name: "5", args: args{s: "1234", start: 3}, want: "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstrToEnd(tt.args.s, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstrToEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SubstrToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
