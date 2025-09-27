package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	votes := make(map[string]int)
	var w sync.WaitGroup

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			votes["Кандидат"]++
			mu.Unlock()
			fmt.Printf("Голос %d принят\n", i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(300 * time.Millisecond)
			mu.Lock()
			count := votes["Кандидат"]
			mu.Unlock()
			fmt.Printf("Текущий счет: %d голосов\n", count)
		}
	}()
	
	w.Wait()
	fmt.Printf("результ: %d голосов\n", votes["Кандидат"])
}
