package main

import (
	"reflect"
	"testing"
)

func TestCreateRandomIntSlice(t *testing.T) {
	randomSlice := createRandomIntSlice()

	if len(randomSlice) != 10 {
		t.Error("Ожидалась длина, равная 10, получена длина ", len(randomSlice))
	}

	for _, v := range randomSlice {
		if v < 0 || v > 99 {
			t.Error("Ожидалось числа в диапазоне от 0 до 99, получено ", v)
		}
	}
}

func TestSliceExample(t *testing.T) {
	randomSlice := createRandomIntSlice()
	evenSlice := sliceExample(randomSlice)

	for _, v := range evenSlice {
		if v%2 != 0 {
			t.Error("Ожидалось числа, которые являются четными, получено ", v)
		}
	}
}

// Тест для функции addElements
func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		element  int
		expected []int
	}{
		{"Добавление в пустой слайс", []int{}, 1, []int{1}},
		{"Добавление в непустой слайс", []int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, tt.element)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}

// Тест для функции copySlice
func TestCopySlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Копирование пустого слайса", []int{}, []int{}},
		{"Копирование непустого слайса", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
			// Проверка, что изменения в копии не влияют на оригинал
			if len(result) > 0 {
				result[0]++
				if reflect.DeepEqual(result, tt.input) {
					t.Errorf("Изменения в копии изменили оригинал: %v", tt.input)
				}
			}
		})
	}
}

// Тест для функции removeElement
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{"Удаление из начала", []int{1, 2, 3}, 0, []int{2, 3}},
		{"Удаление из середины", []int{1, 2, 3}, 1, []int{1, 3}},
		{"Удаление из конца", []int{1, 2, 3}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.input, tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}
