package sorts_test

import (
	"math/rand"
	"testing"

	. "github.com/deanveloper/sorts"
)

var quickResult []int

func TestQuickSort(t *testing.T) {
	x := rand.Perm(50)
	QuickSort(x)
	for i := range x {
		if i != x[i] {
			t.Log(x)
			t.Fatal("sort does not work")
		}
	}
}

func TestParallelQuickSort(t *testing.T) {
	x := rand.Perm(2000)
	ParallelQuickSort(x, nil)
	for i := range x {
		if i != x[i] {
			t.Fatal("sort does not work: expected", i, "but got", x[i])
		}
	}
}

func BenchmarkQuickSortSmall(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(100)
		QuickSort(x)
		r = x
	}
	quickResult = r
}

func BenchmarkParallelQuickSortSmall(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(100)
		ParallelQuickSort(x, nil)
		r = x
	}
	quickResult = r
}

func BenchmarkQuickSortMed(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(2000)
		QuickSort(x)
		r = x
	}
	quickResult = r
}

func BenchmarkParallelQuickSortMed(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(2000)
		ParallelQuickSort(x, nil)
		r = x
	}
	quickResult = r
}

func BenchmarkQuickSortLarge(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(500000)
		QuickSort(x)
		r = x
	}
	quickResult = r
}

func BenchmarkParallelQuickSortLarge(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		x := rand.Perm(500000)
		ParallelQuickSort(x, nil)
		r = x
	}
	quickResult = r
}
