package main

import (
	"math"
	"testing"
)

func TestTransformPipeline(t *testing.T) {
	// Входные и ожидаемые данные
	input := []uint8{1, 2, 3, 4, 5}
	expected := []float64{1, 8, 27, 64, 125}

	// Создаём каналы
	uint8Chan := make(chan uint8, len(input))
	float64Chan := make(chan float64, len(input))

	// Заполняем первый
	go func() {
		for _, v := range input {
			uint8Chan <- v
		}
		close(uint8Chan)
	}()

	// Запускаем функцию
	go makeIntSquareChannel(uint8Chan, float64Chan)

	// Проверяем результат
	var results []float64
	for v := range float64Chan {
		results = append(results, v)
	}

	// Сравниваем
	if len(results) != len(expected) {
		t.Fatalf("ожидалось %d, получено %d", len(expected), len(results))
	}
	for i, v := range results {
		if math.Abs(v-expected[i]) > 1e-9 {
			t.Errorf("неожиданный результат для индекса %d: получили %.2f, ожидали %.2f", i, v, expected[i])
		}
	}
}
