package concurrency_basics_3_conditional_variables

import (
	"math/rand"
	"sync"
	"time"
)

func MutexLimitation()  {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex

	for i:=0; i < 10; i++ {
		go func() {
			vote := requestVote() //sending parallel requests
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
		}()
	}

	for {

		mu.Lock()
		if count >= 5 && finished == 10 {
			break
		} else {
			println("waiting :(")
		}
		mu.Unlock()
		//constant looping consumes CPU cycles
		//time.Sleep(50 * time.Microsecond) // with this it's slightly better, yet not good
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}

func requestVote() bool {
	return true
}