package main

import "fmt"

func main() {
	// Пример работы функции
	sliceA := []string{"apple", "banana", "cherry", "date", "elderberry"}
	sliceB := []string{"cherry", "date", "fig", "grape", "kiwi"}
	fmt.Println("Слайсы перед поиском общих элементов")
	fmt.Println("Slice A:", sliceA)
	fmt.Println("Slice B:", sliceB)

	commonElements := findCommonElementsInSlice(sliceA, sliceB)
	fmt.Println("\nОбщие элементы в слайсах:")
	fmt.Println(commonElements)
}

func findCommonElementsInSlice(a, b []string) []string {
	var maxCapacity int
	if len(a) > len(b) {
		maxCapacity = len(a)
	} else {
		maxCapacity = len(b)
	}
	res := make([]string, 0, maxCapacity)
	for _, v := range a {
		if contains(b, v) {
			res = append(res, v)
		}
	}
	return res
}

func contains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
