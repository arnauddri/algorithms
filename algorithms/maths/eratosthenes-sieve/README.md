# Implementation of the Sieve of Eratosthenes algorithm.

### Sieve of Eratosthenes Overview:

It is a simple, ancient algorithm for finding all prime numbers
up to any given limit. It does so by iteratively marking as composite (i.e. not prime)
the multiples of each prime, starting with the multiples of 2.
The sieve of Eratosthenes is one of the most efficient ways
to find all of the smaller primes (below 10 million or so).

**Time Complexity:** O(n log log n)

Pseudocode: [https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)

### Benchmark

Executed in ```44.334s```:

```
BenchmarkGetAllPrimesTo10                5000000               496 ns/op
BenchmarkGetAllPrimesTo100               1000000              1196 ns/op
BenchmarkGetAllPrimesTo1000               200000              7724 ns/op
BenchmarkGetAllPrimesTo10000               50000             58489 ns/op
BenchmarkGetAllPrimesTo100000               5000            514589 ns/op
BenchmarkGetAllPrimesTo1000000               500           5590372 ns/op
BenchmarkGetAllPrimesTo10000000               20          96922714 ns/op
BenchmarkGetAllPrimesTo100000000               1        1431869649 ns/op
BenchmarkGetAllPrimesTo1000000000              1       24368113980 ns/op
```
