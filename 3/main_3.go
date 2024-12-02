package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

// Конструктор создания экземпляра нашей структуры
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

// Теперь описываем методы
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Get(key string) (int, error) {
	value, exists := m.data[key]
	if !exists {
		return 0, fmt.Errorf("key '%s' not found in map", key)
	}
	return value, nil
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for key, value := range m.data {
		copyMap[key] = value
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}
