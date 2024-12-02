package main

import (
	"reflect"
	"testing"
)

// Тест для findCommonElementsInSlice
func TestFindCommonElementsInSlice(t *testing.T) {
	tests := []struct {
		name     string
		sliceA   []string
		sliceB   []string
		expected []string
	}{
		{"Общие элементы", []string{"apple", "banana", "cherry"}, []string{"banana", "cherry", "date"}, []string{"banana", "cherry"}},
		{"Нет общих элементов", []string{"apple", "banana"}, []string{"date", "fig"}, []string{}},
		{"Один общий элемент", []string{"apple", "banana"}, []string{"banana", "grape"}, []string{"banana"}},
		{"Одинаковые слайсы", []string{"apple", "banana"}, []string{"apple", "banana"}, []string{"apple", "banana"}},
		{"Пустой первый слайс", []string{}, []string{"apple", "banana"}, []string{}},
		{"Пустой второй слайс", []string{"apple", "banana"}, []string{}, []string{}},
		{"Оба слайса пустые", []string{}, []string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findCommonElementsInSlice(tt.sliceA, tt.sliceB)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}

// Тест для contains
func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		target   string
		expected bool
	}{
		{"Элемент найден", []string{"apple", "banana", "cherry"}, "banana", true},
		{"Элемент не найден", []string{"apple", "banana", "cherry"}, "fig", false},
		{"Пустой слайс", []string{}, "banana", false},
		{"Один элемент найден", []string{"banana"}, "banana", true},
		{"Один элемент не найден", []string{"banana"}, "apple", false},
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
