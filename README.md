# sorts

Just a bunch of sorts written in Go. This is for educational purposes, use the [`"sort"`](https://godoc.org/sort) package if you actually want to sort your data quickly.

<details><summary>Benchmarks as of 41fb03</summary>
<pre>
$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/deanveloper/sorts
BenchmarkMergeSortSmall-4                 100000             15214 ns/op
BenchmarkParallelMergeSortSmall-4         100000             15228 ns/op
BenchmarkMergeSortMed-4                     5000            363972 ns/op
BenchmarkParallelMergeSortMed-4             5000            243931 ns/op
BenchmarkMergeSortLarge-4                     10         128263560 ns/op
BenchmarkParallelMergeSortLarge-4             20          75987090 ns/op
BenchmarkQuickSortSmall-4                 200000              7374 ns/op
BenchmarkParallelQuickSortSmall-4         200000              7431 ns/op
BenchmarkQuickSortMed-4                    10000            193740 ns/op
BenchmarkParallelQuickSortMed-4            10000            145030 ns/op
BenchmarkQuickSortLarge-4                     20          69142310 ns/op
BenchmarkParallelQuickSortLarge-4             30          40572696 ns/op
PASS
ok      github.com/deanveloper/sorts    35.831s
</pre></details>