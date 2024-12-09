package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	t.Run("Базовый случай", func(t *testing.T) {
		wg := CustomWaitGroup{sem: make(chan struct{}, 100)}

		wg.Add(3)

		go func() {
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}()
		go func() {
			time.Sleep(200 * time.Millisecond)
			wg.Done()
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			wg.Done()
		}()

		wg.Wait() // Должно завершиться после 300 мс
	})

	t.Run("Паника из-за негативного счётчика", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Ожидалась паника, но её не случилось")
			}
		}()

		wg := CustomWaitGroup{sem: make(chan struct{}, 100)}
		wg.Add(1)
		wg.Done()
		wg.Done() // Вызовет панику, потому что Add был 1 раз, а Done 2
	})

	t.Run("Корректная работа Add и Done", func(t *testing.T) {
		wg := CustomWaitGroup{sem: make(chan struct{}, 100)}

		wg.Add(5)

		go func() {
			for i := 0; i < 5; i++ {
				wg.Done()
			}
		}()

		wg.Wait()
	})
}
