package concurrency_basics_2_mutex

import "time"
import "sync"

func FlakyCounter()  {
	counter := 0
	for i := 0; i < 1000 ;i++ {
		go func() {
			counter = counter + 1
		}()
	}

	time.Sleep(1 * time.Second)
	println(counter)
}

func CounterWithLock()  {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000 ;i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter = counter + 1
		}()
	}

	time.Sleep(1 * time.Second) // ideally waitGroup should be used here.  This example is to showcase mutex.
	mu.Lock()
	defer mu.Unlock()
	println(counter)
}
