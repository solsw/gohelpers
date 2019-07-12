package oshelper

import (
	"testing"

	"github.com/solsw/gohelpers/ioutilhelper"
)

func TestFileExists(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "1e", args: args{fileName: ""}, want: false, wantErr: true},
		{name: "1", args: args{fileName: ioutilhelper.TempFileName0()}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileExists(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
