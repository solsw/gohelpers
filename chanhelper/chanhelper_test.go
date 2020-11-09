package chanhelper

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge2Ints(t *testing.T) {
	var (
		in1 chan int = make(chan int)
		in2 chan int = make(chan int)
	)

	go func() {
		for i := 0; i < 4; i++ {
			in1 <- 1
		}
		close(in1)
	}()

	go func() {
		for i := 0; i < 4; i++ {
			in2 <- 2
		}
		close(in2)
	}()

	var out []int
	for v := range Merge2Ints(in1, in2) {
		out = append(out, v)
	}
	sort.Ints(out)
	want := []int{1, 1, 1, 1, 2, 2, 2, 2}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("MergeInts() = %v, want %v", out, want)
	}
}
