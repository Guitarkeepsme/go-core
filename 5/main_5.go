package main

import "fmt"

func main() {
	// Пример работы функции
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{3, 4, 5, 6, 7}

	isSame, sameElements := NewSliceFromSameElements(sliceA, sliceB)

	fmt.Printf("Результат проверки слайсов: %v, получившийся слайс: %v\n", isSame, sameElements)
}

func NewSliceFromSameElements(a, b []int) (bool, []int) {
	var maxCapacity int
	if len(a) > len(b) {
		maxCapacity = len(a)
	} else {
		maxCapacity = len(b)
	}

	res := make([]int, 0, maxCapacity)

	for _, v := range a {
		if contains(b, v) {
			res = append(res, v)
		}
	}

	return len(res) > 0, res
}

func contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
