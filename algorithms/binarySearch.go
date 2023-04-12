package algorithms

import "log"

func BinarySearch(values []int, toFind int) int {
	if len(values) == 0 {
		return -1
	}
	beginIndex := 0
	endIndex := len(values) - 1

	for beginIndex <= endIndex {
		midIndex := beginIndex + (endIndex-beginIndex)/2
		log.Printf("midIndex: %v, beginIndex: %v, endIndex: %v", midIndex, beginIndex, endIndex)
		mid := values[midIndex]
		if mid == toFind {
			return midIndex
		} else if toFind < mid {
			endIndex = midIndex - 1
		} else if mid < toFind {
			beginIndex = midIndex + 1
		}
	}
	return -1
}

func BinarySearchRec(a []int, search int) (result int, searchCount int) {
	if len(a) == 0 {
		result = -1 // not found
		return
	}
	mid := len(a) / 2
	switch {
	case a[mid] > search:
		result, searchCount = BinarySearchRec(a[:mid], search)
	case a[mid] < search:
		result, searchCount = BinarySearchRec(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}
