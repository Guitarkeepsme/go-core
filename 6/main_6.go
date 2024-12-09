/*
Напишите генератор случайных чисел используя небуфферизированный канал.

Напишите unit тесты к созданным функциям
*/

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// Создаём небуферизированный канал чисел
	chanInt := make(chan int)

	go randomNumberGenerator(chanInt)

	fmt.Println("Открываем канал")
	for i := 0; i < 10; i++ {
		fmt.Println(<-chanInt)
	}
	fmt.Println("Закрываем канал")

}

// Создание генератора чисел
func randomNumberGenerator(ch chan int) {
	// В бесконечном цикле отправляем числа
	for {
		ch <- rand.IntN(1000)
		time.Sleep(time.Second) // Ждём 1 секунду между генерацией чисел
	}
}
