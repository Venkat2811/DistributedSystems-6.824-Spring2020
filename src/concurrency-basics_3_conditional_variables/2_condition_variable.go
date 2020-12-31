package concurrency_basics_3_conditional_variables

import (
	"math/rand"
	"sync"
	"time"
)

func ConditionVarDemo()  {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i:=0; i < 10; i++ {
		go func() {
			vote := requestVote() //sending parallel requests
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			cond.Broadcast()
		}()
	}


	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait() // releases lock, waits to hear from notify or broadcast
		//once it wakes up, lock is acquired again and then condition is checked
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}

// pattern to avoid Lost-Wakeup problem
//
// goroutines doing something on shared var must
// mu.Lock() --> acquire lock first
// --> do something
// cond.BroadCast()
// mu.UnLock()

// waiting threads MUST
// 1 acquire lock mu.Lock()
// while condition = false {
// 		 cond.Wait()
// }
//now condition is true and we have the lock
// mu.UnLock()