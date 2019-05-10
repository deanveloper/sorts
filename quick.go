package sorts

import "sync"

func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	p := partition(slice)

	QuickSort(slice[:p])
	QuickSort(slice[p:])
}

func ParallelQuickSort(slice []int, parentWg *sync.WaitGroup) {
	if len(slice) <= 1 {
		if parentWg != nil {
			parentWg.Done()
		}
		return
	}

	p := partition(slice)

	if len(slice) < 500 {
		ParallelQuickSort(slice[:p], nil)
		ParallelQuickSort(slice[p:], nil)
	} else {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go ParallelQuickSort(slice[:p], wg)
		go ParallelQuickSort(slice[p:], wg)
		wg.Wait()
	}
	if parentWg != nil {
		parentWg.Done()
	}
}

// partition sorts the array into a left and right slice,
// returns the index of the pivot (which is where they are)
// divided
func partition(slice []int) int {
	if len(slice) <= 1 {
		return len(slice) - 1
	}

	left, right := 0, len(slice)-1

	for i := range slice {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	return left
}
