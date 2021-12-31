# sync

### Mutex package

main functions:

1. `Lock`
2. `Unlock`
3. `TryLock`

This package is implemented using channels in golang.
In this package, unlike the official sync package of the golang,
we have semaphores. Also, we have another function that does not 
exist in golang sync package, i.e., `TryLock`

### Single package

It just has one main function:
1. `Do`

### Waitgroup package

main functions:
1. `Add`
2. `Done`
3. `Wait`



