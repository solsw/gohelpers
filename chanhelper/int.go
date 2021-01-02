package chanhelper

import (
	"sync"
)

// IntToInterface converts int channel into interface{} channel.
func IntToInterface(in <-chan int) <-chan interface{} {
	out := make(chan interface{})

	go func(i <-chan int, o chan<- interface{}) {
		defer close(o)
		for v := range i {
			o <- v
		}
	}(in, out)

	return out
}

// PeekInt checks int channel for a value.
// Returns: 'v' - the value (if 'ok' is true and 'open' is true),
// 'ok' - 'true' if the value was obtained from the channel or 'false' if there is no value available,
// 'open' - whether the channel is opened or closed.
func PeekInt(ch <-chan int) (v int, ok, open bool) {
	select {
	case v, open = <-ch:
		ok = open
	default:
		ok = false
		open = true
	}
	return
}

func merge2Int(i1, i2 <-chan int, o chan<- int) {
	defer close(o)
	for i1 != nil || i2 != nil {
		select {
		case v1, ok1 := <-i1:
			if ok1 {
				o <- v1
			} else {
				i1 = nil
			}
		case v2, ok2 := <-i2:
			if ok2 {
				o <- v2
			} else {
				i2 = nil
			}
		}
	}
}

// Merge2Int merges two int channels into one.
func Merge2Int(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go merge2Int(in1, in2, out)
	return out
}

func merge3Int(i1, i2, i3 <-chan int, o chan<- int) {
	defer close(o)
	for i1 != nil || i2 != nil || i3 != nil {
		select {
		case v1, ok1 := <-i1:
			if ok1 {
				o <- v1
			} else {
				i1 = nil
			}
		case v2, ok2 := <-i2:
			if ok2 {
				o <- v2
			} else {
				i2 = nil
			}
		case v3, ok3 := <-i3:
			if ok3 {
				o <- v3
			} else {
				i3 = nil
			}
		}
	}
}

// Merge3Int merges three int channels into one.
func Merge3Int(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	go merge3Int(in1, in2, in3, out)
	return out
}

func merge4Int(i1, i2, i3, i4 <-chan int, o chan<- int) {
	defer close(o)
	for i1 != nil || i2 != nil || i3 != nil || i4 != nil {
		select {
		case v1, ok1 := <-i1:
			if ok1 {
				o <- v1
			} else {
				i1 = nil
			}
		case v2, ok2 := <-i2:
			if ok2 {
				o <- v2
			} else {
				i2 = nil
			}
		case v3, ok3 := <-i3:
			if ok3 {
				o <- v3
			} else {
				i3 = nil
			}
		case v4, ok4 := <-i4:
			if ok4 {
				o <- v4
			} else {
				i4 = nil
			}
		}
	}
}

// Merge4Int merges four int channels into one.
func Merge4Int(in1, in2, in3, in4 <-chan int) <-chan int {
	out := make(chan int)
	go merge4Int(in1, in2, in3, in4, out)
	return out
}

// MergeInt merges several int channels into one.
func MergeInt(ins ...chan int) <-chan int {
	switch len(ins) {
	case 0:
		out := make(chan int)
		close(out)
		return out
	case 1:
		return ins[0]
	case 2:
		return Merge2Int(ins[0], ins[1])
	case 3:
		return Merge3Int(ins[0], ins[1], ins[2])
	case 4:
		return Merge4Int(ins[0], ins[1], ins[2], ins[3])
	default:
		h := len(ins) / 2
		return Merge2Int(MergeInt(ins[:h]...), MergeInt(ins[h:]...))
	}
}

// MergeBufInt merges several int channels into one with buffering values from input channels
// (see https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement/32381409#32381409).
func MergeBufInt(ins ...chan int) <-chan int {
	// https://play.golang.org/p/v4M14CL7OL
	out := make(chan int)
	if len(ins) == 0 {
		close(out)
		return out
	}
	var wg sync.WaitGroup
	wg.Add(len(ins))

	for _, in := range ins {
		go func(i <-chan int, o chan<- int) {
			for v := range i {
				o <- v
			}
			wg.Done()
		}(in, out)
	}

	go func(o chan int) {
		wg.Wait()
		close(o)
	}(out)

	return out
}
