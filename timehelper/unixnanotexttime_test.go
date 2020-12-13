package timehelper

import (
	"reflect"
	"testing"
)

func TestUnixNanoTextTime_MarshalText(t *testing.T) {
	tests := []struct {
		name string
		t    UnixNanoTextTime
		want string
	}{
		{name: "1", t: UnixNanoTextTime(YMDMust(2006, 1, 2)), want: "1136160000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotbytes, _ := tt.t.MarshalText()
			got := string(gotbytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnixNanoTextTime.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixNanoTextTimeUnmarshalStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    UnixNanoTextTime
		wantErr bool
	}{
		{name: "0", args: args{str: "qwerty"}, wantErr: true},
		{name: "1", args: args{str: "1136160000000000000"}, want: UnixNanoTextTime(YMDMust(2006, 1, 2))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unixNanoTextTimeUnmarshalStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("unixNanoTextTimeUnmarshalStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixNanoTextTimeUnmarshalStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
