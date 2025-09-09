package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer wg.Done()
		for u := 1; u <= 5; u++{
			fmt.Println(u)
			time.Sleep(1 * time.Second)//интервал
		}
	}()
	
	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
