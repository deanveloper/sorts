package sorts

import (
	"sync"
)

// MergeSort splits a slice into a left and right slice, then
// continues splitting those two slices recursively until the slices
// are made of single elements. Then it merges all of these slices
// so that all resulting merged slices are sorted from least to greatest.
// When the final two slices are merged, the entire slice is then sorted.
func MergeSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	var left, right []int
	left = append(left, slice[:len(slice)/2]...)
	right = append(right, slice[len(slice)/2:]...)

	MergeSort(left)
	MergeSort(right)

	var l, r int
	for i := range slice {
		if r >= len(right) || (l < len(left) && left[l] < right[r]) {
			slice[i] = left[l]
			l++
		} else {
			slice[i] = right[r]
			r++
		}
	}
}

// ParallelMergeSort works the same way as MergeSort, however when
// splitting up the slices, it is able to do so in parallel.
func ParallelMergeSort(slice []int, parentWg *sync.WaitGroup) {
	if len(slice) <= 1 {
		if parentWg != nil {
			parentWg.Done()
		}
		return
	}

	var left, right []int
	left = append(left, slice[:len(slice)/2]...)
	right = append(right, slice[len(slice)/2:]...)

	// don't do it in parallel if the slice is small
	if len(slice) < 500 {
		ParallelMergeSort(left, nil)
		ParallelMergeSort(right, nil)
	} else {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go ParallelMergeSort(left, wg)
		go ParallelMergeSort(right, wg)
		wg.Wait()
	}

	var l, r int
	for i := range slice {
		if r >= len(right) || (l < len(left) && left[l] < right[r]) {
			slice[i] = left[l]
			l++
		} else {
			slice[i] = right[r]
			r++
		}
	}
	if parentWg != nil {
		parentWg.Done()
	}
}
