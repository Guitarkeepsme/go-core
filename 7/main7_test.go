package main

import (
	"reflect"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name     string
		channels []<-chan any
		expected []any
	}{
		{
			name: "Два канала чисел",
			channels: []<-chan any{
				intChannel([]int{3, 4, 5}),
				intChannel([]int{10, 12}),
			},
			expected: []any{3, 4, 5, 10, 12}},

		{
			name: "Пустые каналы",
			channels: []<-chan any{
				emptyChannel(),
				emptyChannel(),
			},
			expected: []any{},
		},
		{
			name: "Каналы разных типов",
			channels: []<-chan any{
				intChannel([]int{1, 2}),
				stringChannel([]string{"a", "b"}),
			},
			expected: []any{1, 2, "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeChannels(tt.channels...)
			received := []any{}
			for v := range result {
				received = append(received, v)
			}
			if !reflect.DeepEqual(received, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, received)
			}
		})
	}
}

// Вспомогательные функции для testmerge -- создаём каналы разных типов
func intChannel(data []int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for _, v := range data {
			ch <- v
		}
	}()
	return ch
}

func stringChannel(data []string) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for _, v := range data {
			ch <- v
		}
	}()
	return ch
}

func emptyChannel() <-chan any {
	ch := make(chan any)
	close(ch) // Закрываем канал сразу, чтобы он был пуст
	return ch
}

// Написал тест на примере числового канала, но не придумал, как
// сделать его универсальным
func TestWrapChannel(t *testing.T) {
	tests := []struct {
		name     string
		input    func() <-chan int // Создание числового канала
		expected []any             // Ожидаемые данные
	}{
		{
			name: "Числовой канал",
			input: func() <-chan int {
				ch := make(chan int, 4)
				ch <- 1
				ch <- 2
				ch <- 3
				ch <- 4
				close(ch)
				return ch
			},
			expected: []any{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Оборачиваем канал конкретного типа
			wrappedChan := wrapChannel(tt.input())

			// Считываем данные из обёрнутого канала
			result := make([]any, 0, len(wrappedChan))
			for v := range wrappedChan {
				result = append(result, v)
			}

			// Проверяем, совпадают ли результаты
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("wrapChannel() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestGenerateRandomChannels(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int // Ожидаемое количество каналов
	}{
		{"Нет каналов", 0, 0},
		{"Несколько каналов", 5, 5},
		{"Много каналов", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Генерация каналов
			channels := generateRandomChannels(tt.n)

			// Проверка количества каналов
			if len(channels) != tt.expected {
				t.Fatalf("Ожидалось %d каналов, получено %d", tt.expected, len(channels))
			}

			// Проверка содержимого каналов
			for i, ch := range channels {
				data := make([]any, 0)
				for v := range ch { // Читаем данные из канала
					data = append(data, v)
				}

				// Проверка, что данные не пустые
				if len(data) == 0 {
					t.Errorf("Канал %d пуст", i)
				}

				// Проверка типов данных
				for _, v := range data {
					switch v.(type) {
					case int, string, bool, float64:
						// Допустимые типы, всё хорошо
					default:
						t.Errorf("Неожиданный тип данных %T в канале %d", v, i)
					}
				}
			}
		})
	}
}
