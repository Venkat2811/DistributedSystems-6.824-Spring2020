## DistributedSystems-6.824-Spring2020
Distributed Systems pet project learning using
- https://pdos.csail.mit.edu/6.824/index.html 
- https://www.youtube.com/playlist?list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB 

## Exercises
- [] Lab 1: Map-Reduce
- [] Lab 2: Raft-FaultTolerance
- [] Lab 3: KV-Server-With-Raft
- [] Lab 4: KV-Server-With-Shard-And-Raft
    - [] Challenge 1: Garbage collection of state
    - [] challenge 2: Client requests during configuration changes

### Setup

Install go
```
$ brew install go@1.13
```

Setup upstream url to
```
$ git remote add upstream git://g.csail.mit.edu/6.824-golabs-2020
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