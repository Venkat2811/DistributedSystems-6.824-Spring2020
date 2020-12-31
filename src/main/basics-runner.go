package main

// main program to run files in concurrency-basics_*

import (
	//"concurrency-basics_1_closure"
	//"concurrency-basics_2_mutex"
	// "concurrency-basics_2_mutex"
	// "concurrency-basics_2_mutex"
	"concurrency-basics_3_conditional_variables"
)

func main() {
	//concurrency_basics_1_closure.SimpleClosure()
	//concurrency_basics_1_closure.BadClosureLoop()
	//concurrency_basics_1_closure.CorrectClosureLoop()

	//concurrency_basics_2_mutex.StartSleepTimer()
	//concurrency_basics_2_mutex.StartMutexSleepTimer()
	//concurrency_basics_2_mutex.FlakyCounter()
	//concurrency_basics_2_mutex.CounterWithLock()
	//concurrency_basics_2_mutex.TransactionsUsingBadLock()
	//concurrency_basics_2_mutex.TransactionsUsingCorrectLock()
	//concurrency_basics_3_conditional_variables.MutexLimitation()
	concurrency_basics_3_conditional_variables.ConditionVarDemo()
}
