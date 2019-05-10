package sorts_test

import (
	"math/rand"
	"testing"

	. "github.com/deanveloper/sorts"
)

var mergeResult []int

func TestMergeSort(t *testing.T) {
	x := rand.Perm(2000)
	MergeSort(x)
	for i := range x {
		if i != x[i] {
			t.Fatal("sort does not work: expected", i, "but got", x[i])
		}
	}
}

func TestParallelMergeSort(t *testing.T) {
	x := rand.Perm(2000)
	ParallelMergeSort(x, nil)
	for i := range x {
		if i != x[i] {
			t.Fatal("sort does not work: expected", i, "but got", x[i])
		}
	}
}

func BenchmarkMergeSortSmall(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(100)
		MergeSort(x)
		r = x
	}
	mergeResult = r
}

func BenchmarkParallelMergeSortSmall(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(100)
		ParallelMergeSort(x, nil)
		r = x
	}
	mergeResult = r
}

func BenchmarkMergeSortMed(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(2000)
		MergeSort(x)
		r = x
	}
	mergeResult = r
}

func BenchmarkParallelMergeSortMed(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(2000)
		ParallelMergeSort(x, nil)
		r = x
	}
	mergeResult = r
}

func BenchmarkMergeSortLarge(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(500000)
		MergeSort(x)
		r = x
	}
	mergeResult = r
}

func BenchmarkParallelMergeSortLarge(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(500000)
		ParallelMergeSort(x, nil)
		r = x
	}
	mergeResult = r
}
