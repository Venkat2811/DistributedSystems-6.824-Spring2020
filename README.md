## DistributedSystems-6.824-Spring2023
Distributed Systems pet project learning using MIT OpenCourseWare
- https://pdos.csail.mit.edu/6.824/index.html 
- https://www.youtube.com/playlist?list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB 

## Exercises
- [X] Lab 0: Go-Concurrency-Basics
- [ ] Lab 1: Map-Reduce
- [ ] Lab 2: Raft-FaultTolerance
- [ ] Lab 3: KV-Server-With-Raft
- [ ] Lab 4: KV-Server-With-Shard-And-Raft
    - [ ] Challenge 1:
    - [ ] Challenge 2: 
- [ ] Lab 5: Paxos [from 2015 labs](http://nil.csail.mit.edu/6.824/2015/labs/lab-3.html) (if possible)   

## Go Concurrency Basics
Go concurrency primitives exercises based on 1st half of the [lecture](https://www.youtube.com/watch?v=UzzcUS2OHqo&list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB&index=5)

packages for the exercises are named in `concurrency-basics_<>_<>` format

#### Run
To run programs here, use:
```
$ cd src/main
$ go run basics-runner.go
```

### Setup

Install go
```
$ brew install go@1.19
```

Setup upstream url to
```
$ git remote add upstream git://g.csail.mit.edu/6.5840-golabs-2023
```

### Tests

For Mac OS X, install timeout using:
```
https://stackoverflow.com/a/21118126
```

run
```
$ go test path/to/test/file.go
```

### Some Notes

Ignore messages like these from Go RPC package
```
2020/12/30 18:47:47 rpc.Register: method "Done" has 1 input parameters; needs exactly three
```
