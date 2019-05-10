package sorts

import (
	"sync"
)

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
