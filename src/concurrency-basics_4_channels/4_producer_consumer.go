package concurrency_basics_4_channels

import (
	"math/rand"
	"time"
)

func ProducerConsumer()  {
	c := make(chan int)
	for i:= 0; i<4; i++ {
		go doWork(c)
	}

	for {
		v := <-c
		println(v)
	}

}

func doWork(c chan int)  {
	for { //infinite loop
		time.Sleep(100 * time.Millisecond)
		c <- rand.Int()
	}
}