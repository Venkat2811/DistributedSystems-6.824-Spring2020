package concurrency_basics_1_closure

import "sync"

func BadClosureLoop() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			sendRPC(i) //i declared outside goroutine could've been modified already before this is executed
			wg.Done()
		}()
		wg.Wait()
	}
}

func sendRPC(i int)  {
	println(i)
}

func CorrectClosureLoop()  {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			sendRPC(j)
			wg.Done()
		}(i)
		wg.Wait()
	}
}
