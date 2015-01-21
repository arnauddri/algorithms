# Extended Euclidean Algorithm

This module implements the extended greatest common divider algorithm.

*Pre:* two integers a and b

*Post:* a tuple ```(x, y)``` where ```a*x + b*y = gcd(a, b)``` (bezout coefficients)

Pseudo Code: [http://en.wikipedia.org/wiki/Extended_Euclidean_algorithm](http://en.wikipedia.org/wiki/Extended_Euclidean_algorithm)

### Benchmark

for ```a :=131313131```, ```b := 121212121```, the benchmark ran in 6.338s:
```
BenchmarkGetCoeff        1       6255517486 ns/op
```
