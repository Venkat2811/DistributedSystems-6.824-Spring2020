package concurrency_basics_4_channels

func UnBufferedDeadLock()  {
	// uncomment below section, go-runtime cannot detect deadlock
	//go func() {
	//	for {
	//
	//	}
	//}()
	c := make(chan bool)
	<-c
	c <- true
	//or
	// c <- true
	// <-c
	// order doesn't matter
	// there MUST be another goroutine doing the opposite action at the same time, if not, DEADLOCK
}