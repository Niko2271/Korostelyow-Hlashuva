package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()

	fmt.Println("Haчинаем. Нажмите <Enter> для отказа.")
	tick := time.Tick(1 * time.Second)
	for countdown := 1; countdown > 0; countdown++ {
		fmt.Println("f")

		select {
		case <-tick:
		case <-done:
			fmt.Println("остановачка!")
			return
		}
	}
}
