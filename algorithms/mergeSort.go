package algorithms

import "log"

func MergeSort(values []int) []int {
	log.Printf("MergeSort: %+v", values)
	lenOfValues := len(values)
	if lenOfValues == 1 {
		return values
	}
	lenOfHalfValues := lenOfValues / 2
	firstHalf := values[:lenOfHalfValues]
	secondHalf := values[lenOfHalfValues:]
	return merge(MergeSort(firstHalf), MergeSort(secondHalf))
}

func merge(first, second []int) []int {
	log.Printf("MERGE %+v %+v", first, second)
	res := make([]int, len(first)+len(second))

	i := 0
	for len(first) > 0 && len(second) > 0 {
		if first[0] < second[0] {
			res[i] = first[0]
			first = first[1:]
		} else {
			res[i] = second[0]
			second = second[1:]
		}
		i++
	}
	for j := 0; j < len(first); j++ {
		res[i] = first[j]
		i++
	}
	for j := 0; j < len(second); j++ {
		res[i] = second[j]
		i++
	}
	log.Printf("MERGE OUT %+v", res)
	return res
}
