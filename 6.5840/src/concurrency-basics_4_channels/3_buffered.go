package concurrency_basics_4_channels

func BufferedChannel()  {
	c := make(chan bool, 1) //until capacity reached, it works fine,
	// once capacity is reached, behaves the same way as unbuffered
	c <- true
	<-c
}
