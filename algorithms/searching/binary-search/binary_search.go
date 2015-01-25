package bs

import ()

func search(sortedArray []int, el int) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		middle := ((end - init) >> 1) + init

		if sortedArray[middle] == el {
			return middle
		}

		if sortedArray[middle] < el {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}
