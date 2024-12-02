package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func main() {
	randomSlice := createRandomIntSlice()

	randomInt := rand.Intn(500)

	fmt.Println("Случайный слайс:", randomSlice)

	fmt.Printf("Удаляем нечётные элементы из слайса: %v\n", sliceExample(randomSlice))

	fmt.Printf("Добавляем в конец слайса случайное число %v: %v\n",
		randomInt, addElements(randomSlice, randomInt))

	fmt.Printf("Копируем первоначальный слайс: %v\n", copySlice(randomSlice))

	fmt.Printf("Удаляем случайный элемент слайса: %v\n", removeElement(randomSlice, rand.Intn(len(randomSlice))))

}
func createRandomIntSlice() (res []int) {
	for i := 0; i < 10; i++ {
		res = append(res, rand.Intn(100))
	}

	return res
}

func sliceExample(slice []int) (res []int) {
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			res = append(res, slice[i])
		}
	}
	return res
}

func addElements(slice []int, n int) []int {
	// Неужели так просто?..
	slice = append(slice, n)
	return slice
}

func copySlice(original []int) []int {
	clone := make([]int, len(original))

	copy(clone, original)
	return clone
}

func removeElement(slice []int, i int) []int {
	// Удаляем элемент i: возвращаем слайс до i-того элемента + слайс после него
	return append(slice[:i], slice[i+1:]...)
}
