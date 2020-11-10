package chanhelper

import (
	"sync"
)

// Merge2Ints merges two int channels into one.
func Merge2Ints(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func(in1, in2 <-chan int, out chan<- int) {
		var (
			v   int
			ok1 bool = true
			ok2 bool = true
		)
		for ok1 || ok2 {
			select {
			case v, ok1 = <-in1:
				if ok1 {
					out <- v
				}
			case v, ok2 = <-in2:
				if ok2 {
					out <- v
				}
			}
		}
		close(out)
	}(in1, in2, out)
	return out
}

// MergeNInts merges several int channels into one.
func MergeNInts(ins ...chan int) <-chan int {
	if len(ins) == 0 {
		return nil
	}
	// https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement/32381409#32381409
	// https://play.golang.org/p/v4M14CL7OL
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, c := range ins {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
