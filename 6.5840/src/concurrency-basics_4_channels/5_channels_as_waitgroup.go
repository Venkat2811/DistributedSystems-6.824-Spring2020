package concurrency_basics_4_channels

func sendRPC(i int)  {
	println(i)
}

func ClosureLoopWithChannel()  {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(j int) {
			sendRPC(j)
			done <- true
		}(i)
	}

	for i:=0; i<5; i++ {
		<-done
	}
}
