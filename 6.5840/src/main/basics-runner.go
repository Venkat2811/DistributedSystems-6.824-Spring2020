package main

// main program to run files in concurrency-basics_*

import (
	concurrency_basics_1_closure "6.5840/concurrency-basics_1_closure"
	concurrency_basics_2_mutex "6.5840/concurrency-basics_2_mutex"
	concurrency_basics_3_conditional_variables "6.5840/concurrency-basics_3_conditional_variables"
	concurrency_basics_4_channels "6.5840/concurrency-basics_4_channels"
)

func main() {
	concurrency_basics_1_closure.SimpleClosure()
	concurrency_basics_1_closure.BadClosureLoop()
	concurrency_basics_1_closure.CorrectClosureLoop()

	concurrency_basics_2_mutex.StartSleepTimer()
	concurrency_basics_2_mutex.StartMutexSleepTimer()
	concurrency_basics_2_mutex.FlakyCounter()
	concurrency_basics_2_mutex.CounterWithLock()
	concurrency_basics_2_mutex.TransactionsUsingBadLock()
	concurrency_basics_2_mutex.TransactionsUsingCorrectLock()

	concurrency_basics_3_conditional_variables.MutexLimitation()
	concurrency_basics_3_conditional_variables.ConditionVarDemo()

	concurrency_basics_4_channels.UnbufferedChannel()
	concurrency_basics_4_channels.UnBufferedDeadLock()
	concurrency_basics_4_channels.BufferedChannel()
	concurrency_basics_4_channels.ProducerConsumer()
	concurrency_basics_4_channels.ClosureLoopWithChannel()
}
