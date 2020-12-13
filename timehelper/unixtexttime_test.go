package timehelper

import (
	"reflect"
	"testing"
)

func TestUnixTextTime_MarshalText(t *testing.T) {
	tests := []struct {
		name string
		t    UnixTextTime
		want string
	}{
		{name: "1", t: UnixTextTime(YMDMust(2006, 1, 2)), want: "1136160000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotbytes, _ := tt.t.MarshalText()
			got := string(gotbytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnixTextTime.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixTextTimeUnmarshalStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    UnixTextTime
		wantErr bool
	}{
		{name: "0", args: args{str: "qwerty"}, wantErr: true},
		{name: "1", args: args{str: "1136160000"}, want: UnixTextTime(YMDMust(2006, 1, 2))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unixTextTimeUnmarshalStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("unixTextTimeUnmarshalStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixTextTimeUnmarshalStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
