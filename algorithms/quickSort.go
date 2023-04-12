package algorithms

import "math/rand"

func QuickSort(values []int) []int {
	lenOfValues := len(values)
	if lenOfValues <= 1 {
		return values
	}
	pivotIndex := rand.Intn(lenOfValues)
	pivot := values[pivotIndex]
	remove(values, pivotIndex)

	var moreThanPivot, lessThanPivot []int
	for i := 0; i < len(values)-1; i++ {
		if values[i] < pivot {
			lessThanPivot = append(lessThanPivot, values[i])
		} else {
			moreThanPivot = append(moreThanPivot, values[i])
		}
	}
	var res []int
	res = append(QuickSort(lessThanPivot), pivot)
	res = append(res, QuickSort(moreThanPivot)...)
	return res
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
