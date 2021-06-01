package contexthelper

func andDone(done1, done2 <-chan struct{}, done chan<- struct{}) {
	for done1 != nil || done2 != nil {
		select {
		case _, ok1 := <-done1:
			if !ok1 {
				done1 = nil
			}
		case _, ok2 := <-done2:
			if !ok2 {
				done2 = nil
			}
		}
	}
	close(done)
}

func orDone(done1, done2 <-chan struct{}, done chan<- struct{}) {
	// done1 and done2 are not both nil here
	select {
	case <-done1:
	case <-done2:
	}
	close(done)
}
