package chanhelper

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
