# Fibonacci's Sequence

Get the n-th Fibonacci's number by 3 different methods:
* Iterative (O (n))
* Recursive (O (exp n))
* Matrix (O (log n))

Pseudo-code [http://en.wikipedia.org/wiki/Fibonacci_number](http://en.wikipedia.org/wiki/Fibonacci_number)

### Benchmark

The recursive method was not tested due to its computation time.

Passed in ```30,110s```:

```
BenchmarkIter10        100000000       13.3 ns/op
BenchmarkIter100        20000000       109 ns/op
BenchmarkIter1000        2000000       937 ns/op
BenchmarkIter10000        200000      9108 ns/op
BenchmarkIter100000        20000     93367 ns/op
BenchmarkIter1000000        2000   1017469 ns/op
BenchmarkIter10000000        100  10168824 ns/op

BenchmarkMatrix10        1000000      2314 ns/op
BenchmarkMatrix100        500000      3984 ns/op
BenchmarkMatrix1000       500000      8713 ns/op
BenchmarkMatrix10000      200000      9513 ns/op
BenchmarkMatrix100000     100000     13452 ns/op
BenchmarkMatrix1000000    100000     15263 ns/op
BenchmarkMatrix10000000   100000     15554 ns/op
```
