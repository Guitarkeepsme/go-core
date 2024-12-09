/*
	Сделать конвейер чисел

Даны два канала.
В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.

Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

* Напишите unit тесты к созданным функциям
*/
package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	// Создаём каналы
	intChan := make(chan uint8, 10)
	floatChan := make(chan float64, 10)

	// Заполняем каналы числами
	go func() {
		ints := make([]uint8, 0, 10)
		for i := 0; i < 10; i++ {
			ints = append(ints, (uint8(rand.UintN(30))))
		}
		fmt.Printf("Исходные числа: %v\n", ints)
		for _, n := range ints {
			intChan <- n
		}
		close(intChan)
	}()

	// Запускаем конвейер
	go makeIntSquareChannel(intChan, floatChan)

	// Выводим результаты в консоль
	for res := range floatChan {
		fmt.Printf("Результат: %.2f\n", res)
		time.Sleep(time.Millisecond * 300)
	}
}

func makeIntSquareChannel(input <-chan uint8, output chan<- float64) {
	for num := range input {
		res := math.Pow(float64(num), 3)
		output <- res
	}
	close(output)
}
