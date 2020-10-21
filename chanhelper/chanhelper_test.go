package chanhelper

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeInts(t *testing.T) {
	in1 := make(chan int)
	in2 := make(chan int)

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

	var ii []int
	for v := range MergeInts(in1, in2) {
		ii = append(ii, v)
	}
	sort.Ints(ii)
	want := []int{1, 1, 1, 1, 2, 2, 2, 2}
	if !reflect.DeepEqual(ii, want) {
		t.Errorf("MergeInts() = %v, want %v", ii, want)
	}
}
