package concurrency_basics_2_mutex

import "time"
import "sync"

var done = false
var mu sync.Mutex

func StartMutexSleepTimer()  {
	time.Sleep(1 * time.Second)
	println("started")
	go periodicWithMutex()
	time.Sleep(5 * time.Second) //main thread waits
	mu.Lock()
	done = true
	mu.Unlock()
	println("stopped")
	time.Sleep(3 * time.Second)
}

func periodicWithMutex()  {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			mu.Unlock()
			break
		}
		mu.Unlock()
	}
}
