/*
Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

* Напишите unit тесты к созданным функциям
*/

package main

import "sync"

type CustomWaitGroup struct {
	counter int
	mu      sync.Mutex
	sem     chan struct{}
}

// Add увеличивает или уменьшает счётчик задач
func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter += delta
	// Если счётчик стал отрицательным, это ошибка
	if wg.counter < 0 {
		panic("отрицательное значение счётчика в CustomWaitGroup")
	}

	// Если задача добавлена, увеличиваем канал
	for i := 0; i < delta; i++ {
		wg.sem <- struct{}{}
	}
}

// Done уменьшает счётчик задач на 1.
func (wg *CustomWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	// Уменьшаем счётчик
	if wg.counter <= 0 {
		panic("CustomWaitGroup.Done был вызван слишком много раз")
	}

	wg.counter--
	<-wg.sem // Убираем элемент из канала
}

// Wait ожидает завершения всех задач.
func (wg *CustomWaitGroup) Wait() {
	// Ждём, пока канал не опустеет
	for i := 0; i < wg.counter; i++ {
		<-wg.sem
	}
}
