package nullable

import (
	"testing"
)

func TestInt_Set(t *testing.T) {
	int0 := NullInt()
	type args struct {
		v int
	}
	tests := []struct {
		name string
		n    *Int
		args args
	}{
		{name: "0", n: &int0, args: args{v: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Set(tt.args.v)
			if *tt.n.val != tt.args.v {
				t.Errorf("%v, want %v", *tt.n.val, tt.args.v)
			}
		})
	}
}
