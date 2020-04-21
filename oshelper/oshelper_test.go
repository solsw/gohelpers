package oshelper

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	f, _ := ioutil.TempFile("", "")
	f.Close()
	defer os.Remove(f.Name())
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "1e", args: args{filename: ""}, want: false, wantErr: true},
		{name: "1", args: args{filename: "C3043E18D2234F2897BE0BCEBBE0C840"}, want: false},
		{name: "2", args: args{filename: f.Name()}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileExists(tt.args.filename)
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
