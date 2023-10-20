package concurrency_basics_2_mutex

import (
	"sync"
	"time"
	"fmt"
)

func TransactionsUsingBadLock()  {
	alice := 10000
	bob := 10000
	var mu sync.Mutex

	total := alice + bob

	go func() {
		for i:=0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			mu.Unlock()
			mu.Lock()
			bob += 1
			mu.Unlock()
		}
	}()

	go func() {
		for i:=0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			mu.Unlock()
			mu.Lock()
			alice += 1
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("observed violation, alice = %v, bob = %v\n", alice, bob)
		}
		mu.Unlock()
	}

	if alice+bob == total {
		println("final balance same")
	}

}

func TransactionsUsingCorrectLock()  {
	alice := 10000
	bob := 10000
	var mu sync.Mutex

	total := alice + bob

	go func() {
		for i:=0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			bob += 1
			mu.Unlock()
		}
	}()

	go func() {
		for i:=0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			alice += 1
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("observed violation, alice = %v, bob = %v\n", alice, bob)
		}
		mu.Unlock()
	}

	if alice+bob == total {
		println("final balance same")
	}

}

