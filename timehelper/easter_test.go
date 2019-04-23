package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestCatholicEaster(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{year: 2018},
			want: time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC)},
		{name: "2", args: args{year: 2019},
			want: time.Date(2019, time.April, 21, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CatholicEaster(tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CatholicEaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrthodoxEaster(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{year: 2018},
			want: time.Date(2018, time.April, 8, 0, 0, 0, 0, time.UTC)},
		{name: "2", args: args{year: 2019},
			want: time.Date(2019, time.April, 28, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrthodoxEaster(tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrthodoxEaster() = %v, want %v", got, tt.want)
			}
		})
	}
}
