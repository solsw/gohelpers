package timehelper

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTimeString(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    TimeString
		wantErr bool
	}{
		{name: "0", args: args{t: time.Time{}}, want: ""},
		{name: "1", args: args{t: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)}, want: "2006-01-02T15:04:05Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeString(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeString_Time(t *testing.T) {
	tests := []struct {
		name    string
		ts      TimeString
		want    time.Time
		wantErr bool
	}{
		{name: "0", ts: TimeString(""), want: time.Time{}},
		{name: "1", ts: TimeString("2006-01-02T15:04:05Z"), want: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.Time()
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeString.Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeString.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
