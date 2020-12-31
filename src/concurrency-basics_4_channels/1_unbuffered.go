package concurrency_basics_4_channels

import (
	"fmt"
	"time"
)

func UnbufferedChannel()  {

	c := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		<-c //receiving
	}()
	start := time.Now()
	c <- true //sending
	// blocks until other goroutine receives
	fmt.Printf("send took %v\n", time.Since(start))

}