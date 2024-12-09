/*
Напишите программу на Go, которая сливает N каналов в один.

* Напишите unit тесты к созданным функциям
*/

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const minimumAmountOfChannels = 2

func main() {
	// Создаём n каналов, но не меньше 2
	channelsAmount := rand.IntN(10) + minimumAmountOfChannels
	channels := generateRandomChannels(channelsAmount)
	fmt.Printf("Создано каналов: %d\n", channelsAmount)

	time.Sleep(1 * time.Second)
	// Объединяем каналы
	mergedChannels := mergeChannels(channels...)

	fmt.Println("Читаем из общего канала:")
	// Читаем из объединённого канала
	for data := range mergedChannels {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(data)
	}
}

// Функция для слияния каналов, которая принимает N каналов и возвращает один
func mergeChannels(chs ...<-chan any) chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for _, ch := range chs {
			for v := range ch {
				out <- v
			}
		}
	}()
	return out
}

// Оборачиваем любой канал в any
func wrapChannel[T any](in <-chan T) <-chan any {
	out := make(chan any)

	go func() {
		defer close(out)
		for v := range in {
			out <- v
		}
	}()
	return out
}

// Генерируем случайное число каналов
func generateRandomChannels(n int) []<-chan any {
	channels := make([]<-chan any, n)
	for i := 0; i < n; i++ {
		// Создаём один из стандартных типов каналов
		switch rand.IntN(4) {
		case 0:
			ch := make(chan int)
			go func() {
				defer close(ch)
				for j := 0; j < 10; j++ {
					ch <- rand.IntN(50)
				}
			}()
			channels[i] = wrapChannel(ch)
		case 1:
			ch := make(chan string)
			go func() {
				defer close(ch)
				for j := 0; j < 10; j++ {
					ch <- fmt.Sprintf("str-%d", rand.IntN(100))
				}
			}()
			channels[i] = wrapChannel(ch)
		case 2:
			ch := make(chan bool)
			go func() {
				defer close(ch)
				for j := 0; j < 3; j++ {
					ch <- rand.IntN(2) == 1
				}
			}()
			channels[i] = wrapChannel(ch)
		case 3:
			ch := make(chan float64)
			go func() {
				defer close(ch)
				for j := 0; j < 10; j++ {
					ch <- rand.Float64()
				}
			}()
			channels[i] = wrapChannel(ch)
		default:
			panic("Как мы попали сюда?..")
		}
	}
	return channels
}
