package main

import (
	"reflect"
	"testing"
)

// Тест для NewSliceFromSameElements
func TestNewSliceFromSameElements(t *testing.T) {
	tests := []struct {
		name          string
		sliceA        []int
		sliceB        []int
		expectedBool  bool
		expectedSlice []int
	}{
		{"Общие элементы", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}, true, []int{3, 4, 5}},
		{"Нет общих элементов", []int{1, 2}, []int{3, 4}, false, []int{}},
		{"Один общий элемент", []int{1, 2, 3}, []int{3, 6, 7}, true, []int{3}},
		{"Одинаковые слайсы", []int{1, 2, 3}, []int{1, 2, 3}, true, []int{1, 2, 3}},
		{"Пустой первый слайс", []int{}, []int{1, 2, 3}, false, []int{}},
		{"Пустой второй слайс", []int{1, 2, 3}, []int{}, false, []int{}},
		{"Оба слайса пустые", []int{}, []int{}, false, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSame, result := NewSliceFromSameElements(tt.sliceA, tt.sliceB)
			if isSame != tt.expectedBool {
				t.Errorf("Ожидалось %v, получено %v", tt.expectedBool, isSame)
			}
			if !reflect.DeepEqual(result, tt.expectedSlice) {
				t.Errorf("Ожидалось %v, получено %v", tt.expectedSlice, result)
			}
		})
	}
}

// Тест для contains
func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		target   int
		expected bool
	}{
		{"Элемент найден", []int{1, 2, 3, 4, 5}, 3, true},
		{"Элемент не найден", []int{1, 2, 3, 4, 5}, 6, false},
		{"Пустой слайс", []int{}, 3, false},
		{"Один элемент найден", []int{5}, 5, true},
		{"Один элемент не найден", []int{5}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := contains(tt.slice, tt.target)
			if result != tt.expected {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}
