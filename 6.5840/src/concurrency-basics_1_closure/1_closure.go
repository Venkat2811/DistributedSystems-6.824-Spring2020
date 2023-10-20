package concurrency_basics_1_closure

import (
	"sync"
	"time"
)

func usingSleep()  {
	var a string
	go func() {
		a = "hello world"
	}()
	time.Sleep(100 * time.Microsecond)
	println(a)
}

func SimpleClosure()  {
	var a string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "hello world"
		wg.Done()
	}()
	wg.Wait()
	println(a)

}
