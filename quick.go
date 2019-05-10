package sorts

import "sync"

func QuickSort(slice []int) {
	if len(slice) == 1 {
		return
	}

	p := partition(slice)

	QuickSort(slice[:p])
	QuickSort(slice[p:])
}

func ParallelQuickSort(slice []int, parentWg *sync.WaitGroup) {
	if len(slice) == 1 {
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
	var pivot = slice[len(slice)-1]
	var left, right int
	for i, v := range slice {
		if i == len(slice)-1 {
			continue
		}
		if v < pivot {
			swap(&slice[left+1], &slice[i])
			left++
			right++
		} else {
			right++
		}
	}
	swap(&slice[left+1], &slice[len(slice)-1])
	return left + 1
}

func swap(i, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
