package concurrency_basics_2_mutex

import "time"

func StartSleepTimer()  {
	time.Sleep(1 * time.Second)
	println("started")
	go periodic()
	time.Sleep(5 * time.Second) //main thread waits
	//once main exits, all goroutines are killed as well
}

func periodic()  {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
	}
}