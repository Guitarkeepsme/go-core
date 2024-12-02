package main

import (
	"reflect"
	"testing"
)

// Тест для конструктора NewStringIntMap
func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	if m == nil {
		t.Fatal("Конструктор вернул nil")
	}
	if len(m.data) != 0 {
		t.Fatalf("Ожидалось пустая структура, но длина %d", len(m.data))
	}
}

// Тест для метода Add и Exists
func TestStringIntMap_Add_Exists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)

	if !m.Exists("key1") {
		t.Errorf("Ключ 'key1' должен существовать в карте")
	}

	if m.Exists("key2") {
		t.Errorf("Ключ 'key2' не должен существовать в карте")
	}
}

// Тест для метода Get
func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)

	value, err := m.Get("key1")
	if err != nil {
		t.Errorf("Неожиданная ошибка: %v", err)
	}
	if value != 100 {
		t.Errorf("Ожидалось значение 100, но получено %d", value)
	}

	_, err = m.Get("key2")
	if err == nil {
		t.Errorf("Ожидалась ошибка для несуществующего ключа 'key2'")
	}
}

// Тест для метода Remove
func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Remove("key1")

	if m.Exists("key1") {
		t.Errorf("Ключ 'key1' должен быть удалён")
	}

	_, err := m.Get("key1")
	if err == nil {
		t.Errorf("Ожидалась ошибка при получении удалённого ключа")
	}
}

// Тест для метода Copy
func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	copyMap := m.Copy()

	if !reflect.DeepEqual(m.data, copyMap) {
		t.Errorf("Копия карты не совпадает с оригиналом. Оригинал: %v, Копия: %v", m.data, copyMap)
	}

	// Проверяем независимость копии
	copyMap["key1"] = 300
	if m.data["key1"] == copyMap["key1"] {
		t.Errorf("Копия не должна влиять на оригинал")
	}
}
