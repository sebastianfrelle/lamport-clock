Lamport clock implementation in Go, using goroutines to represent processes and buffered Go channels to represent channels between processes.

- Since Go channels are FIFO, this implementation does not capture Lamport clocks' shortcomings with respect to causal ordering.
- The implementation does not model events other than logging that they occurred to stdout.