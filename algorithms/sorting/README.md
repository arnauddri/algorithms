# Sorting

### Heapsort

***Time Complexity:***
* Worst case: O(n log(n))
* Average case: O(n log(n))
* Best case: O(n log(n))

***Space Complexity:***
* O(1)

### Quicksort

**Time Complexity:**
* Worst case: O(n log(n))
* Average case: O(n log(n))
* Best case: O(n^2)

**Space Complexity:**
* O(n)

### Mergesort

**Time Complexity:**
* Best case: O(n log(n))
* Average case: O(n log(n))
* Worst case: O(n log(n))

**Space Complexity:**
* O(n)

### Bubble

**Time Complexity:**
* Best case: O(n)
* Average case: O(n^2)
* Worst case: O(n^2)

**Space Complexity:**
* O(1)

### Insertion

**Time Complexity:**
* Best case: O(n)
* Average case: O(n^n)
* Worst case: O(n^2)

**Space Complexity:**
* O(n)

### Select

**Time Complexity:**
* Best case: O(n^2)
* Average case: O(n^2)
* Worst case: O(n^2)

**Space Complexity:**
* O(1)


### Benchmark Results

```
PASS
BenchmarkBubbleSort100	10000000	       145 ns/op
BenchmarkBubbleSort1000	 2000000	      1272 ns/op
BenchmarkBubbleSort10000	  100000	     14371 ns/op
BenchmarkBubbleSort100000	       1	24769404949 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/bubble-sort	33.273s

PASS
BenchmarkHeapSort100	   20000	     95676 ns/op
BenchmarkHeapSort1000	     500	   3613389 ns/op
BenchmarkHeapSort10000	       5	 287265292 ns/op
BenchmarkHeapSort100000	       1	37184383822 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/heap-sort	43.738s

PASS
BenchmarkInsertionSort100	10000000	       228 ns/op
BenchmarkInsertionSort1000	 1000000	      1857 ns/op
BenchmarkInsertionSort10000	  100000	     18822 ns/op
BenchmarkInsertionSort100000	       1	3949000733 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/insertion-sort	12.348s

PASS
BenchmarkMergeSort100	  100000	     17563 ns/op
BenchmarkMergeSort1000	   10000	    186034 ns/op
BenchmarkMergeSort10000	    1000	   1946963 ns/op
BenchmarkMergeSort100000	      50	  20140692 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/merge-sort	8.193s

PASS
BenchmarkQuickSort100	  500000	      3607 ns/op
BenchmarkQuickSort1000	   50000	     44601 ns/op
BenchmarkQuickSort10000	    5000	    543711 ns/op
BenchmarkQuickSort100000	     500	   6070191 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/quick-sort	10.821s

PASS
BenchmarkSelectionSort100	  500000	      6461 ns/op
BenchmarkSelectionSort1000	    2000	    547985 ns/op
BenchmarkSelectionSort10000	      50	  53036282 ns/op
BenchmarkSelectionSort100000	       1	5344170280 ns/op
ok  	github.com/arnauddri/algorithms/algorithms/sorting/selection-sort	12.653s

PASS
BenchmarkShellSort100	 2000000	       845 ns/op
BenchmarkShellSort1000	  200000	     13049 ns/op
BenchmarkShellSort10000	   10000	    175454 ns/op
BenchmarkShellSort100000	     500	   2485845 ns/op
```
