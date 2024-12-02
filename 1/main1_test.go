package main

import (
	"crypto/sha256"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/exp/rand"
)

func TestGenerateRandomValues(t *testing.T) {
	// Устанавливаем фиксированный сид для генерации предсказуемых случайных чисел
	rand.Seed(42)

	values := generateRandomValues()
	if len(values) != 7 {
		t.Fatalf("Ожидалось 7 значений, получено: %d", len(values))
	}

	// Проверка типов
	expectedTypes := []reflect.Type{
		reflect.TypeOf(0),            // int
		reflect.TypeOf(0),            // int (из другой системы счисления)
		reflect.TypeOf(""),           // string (шестнадцатеричная система)
		reflect.TypeOf(0.0),          // float64
		reflect.TypeOf(""),           // string
		reflect.TypeOf(true),         // bool
		reflect.TypeOf(complex64(0)), // complex64
	}

	for i, v := range values {
		if reflect.TypeOf(v) != expectedTypes[i] {
			t.Errorf("На позиции %d ожидался тип %v, получен: %v", i, expectedTypes[i], reflect.TypeOf(v))
		}
	}
}

func TestAllValuesToString(t *testing.T) {
	values := []interface{}{1, "test", true, 3.14}
	expected := "1 test true 3.14"

	result := allValuesToString(values)
	if result != expected {
		t.Errorf("Ожидалась строка %q, получена: %q", expected, result)
	}
}

func TestHashRunes(t *testing.T) {
	input := []rune("test-input")
	expectedSaltedString := "test-go-2024input"

	hash := hashRunes(input)
	expectedHash := sha256.Sum256([]byte(expectedSaltedString))

	if !reflect.DeepEqual(hash, expectedHash) {
		t.Errorf("Ожидался хэш %x, получен: %x", expectedHash, hash)
	}
}

func TestRandString(t *testing.T) {
	rand.Seed(42) // Фиксируем сид
	result := randString(10)

	if len(result) != 10 {
		t.Errorf("Ожидалась строка длиной 10 символов, получена: %d", len(result))
	}

	// Проверяем, что строка содержит только допустимые символы
	for _, r := range result {
		if !strings.ContainsRune(latinLetters, r) {
			t.Errorf("Найден недопустимый символ: %q", r)
		}
	}
}

func TestRandBoolean(t *testing.T) {
	rand.Seed(42) // Фиксируем сид
	result := randBoolean()

	// Проверка результата: он должен быть либо true, либо false (всегда так, но фиксируем это для читаемости теста)
	if result != true && result != false {
		t.Errorf("Ожидалось значение true или false, получено: %v", result)
	}
}
