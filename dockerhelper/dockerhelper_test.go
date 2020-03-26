package dockerhelper

import "testing"

func TestFromDocker(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "1", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InDocker(); got != tt.want {
				t.Errorf("InDocker() = %v, want %v", got, tt.want)
			}
		})
	}
}
