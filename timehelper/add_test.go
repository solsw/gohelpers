package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMonth(t *testing.T) {
	type args struct {
		t    time.Time
		mons int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{t: DateYMD(1960, time.June, 23), mons: 2}, want: DateYMD(1960, time.August, 23)},
		{name: "2", args: args{t: DateYMD(1960, time.June, 23), mons: 23}, want: DateYMD(1962, time.May, 23)},
		{name: "3", args: args{t: DateYMD(1960, time.June, 23), mons: -2}, want: DateYMD(1960, time.April, 23)},
		{name: "4", args: args{t: DateYMD(1960, time.June, 23), mons: -23}, want: DateYMD(1958, time.July, 23)},
		{name: "5", args: args{t: DateYMD(1960, time.January, 23), mons: -2}, want: DateYMD(1959, time.November, 23)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMonth(tt.args.t, tt.args.mons); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
