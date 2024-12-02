package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand/v2"
	"strconv"
)

const latinLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	// Генерируем случайные переменные и выводим их типы и значения
	randomValues := generateRandomValues()
	for _, v := range randomValues {
		printRandomValueAndType(v)
	}

	// Выводим все значения в одну строку
	allString := allValuesToString(randomValues)

	fmt.Printf("Переводим все значения в строку: %v\n", allString)

	// Теперь переводим полученную строку в слайс рун
	runes := []rune(allString)
	fmt.Printf("Переводим полученную строку в слайс runes: %v\n", runes)

	// Вставляем соль "go-2024" в середину строки
	midIndex := len(runes) / 2
	runesWithSalt := append(runes[:midIndex], append([]rune("go-2024"), runes[midIndex:]...)...)
	fmt.Printf("Строка с добавленной солью: %v\n", string(runesWithSalt))

	hash := hashRunes(runesWithSalt)
	fmt.Printf("SHA-256 хэш: %x\n", hash)

}
func randString(n int) string {
	var letters = []rune(latinLetters)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

// Генерация true/false
func randBoolean() bool {
	return rand.IntN(2) == 1
}

func printRandomValueAndType(v interface{}) {

	// Выводим значения и их типы
	switch v := v.(type) {
	case int:
		fmt.Printf("Тип: int, Значение: %d\n", v)
	case float64:
		fmt.Printf("Тип: float64, Значение: %f\n", v)
	case string:
		fmt.Printf("Тип: string, Значение: %s\n", v)
	case bool:
		fmt.Printf("Тип: bool, Значение: %t\n", v)
	case complex64:
		fmt.Printf("Тип: complex64, Значение: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T, Значение: %v\n", v, v)
	}
}

// Функция для генерации случайных переменных
func generateRandomValues() []interface{} {
	// Генерируем случайное число в десятичной системе
	randomInt10 := rand.IntN(1000)
	// Преобразуем число в другие системы счисления
	randomInt8Str := strconv.FormatInt(int64(randomInt10), 8)
	// Переводим в тип int
	randomInt8, err := strconv.Atoi(randomInt8Str)
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа:", err)
		return nil
	}

	randomInt16 := strconv.FormatInt(int64(randomInt10), 16)

	// Создаём случайное число типа float64
	randomFloat := rand.Float64()

	// Для создания комплексного числа complex64 нужно два float32
	randomFirstFloatForComplex := rand.Float32()
	randomSecondFloatForComplex := rand.Float32()

	randomString := randString(rand.IntN(20))
	randomBoolean := randBoolean()

	complexNumber := complex(randomFirstFloatForComplex, randomSecondFloatForComplex)

	return []interface{}{
		randomInt10,
		randomInt8,
		randomInt16,
		randomFloat,
		randomString,
		randomBoolean,
		complexNumber,
	}
}

func allValuesToString(values []interface{}) string {
	result := ""
	for _, v := range values {
		result += fmt.Sprintf("%v ", v)
	}
	return result[:len(result)-1]
}

func hashRunes(runes []rune) [32]byte {
	midIndex := len(runes) / 2
	runesWithSalt := append(runes[:midIndex], append([]rune("go-2024"), runes[midIndex:]...)...)
	// Хэшируем срез рун с солью
	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hash
}
