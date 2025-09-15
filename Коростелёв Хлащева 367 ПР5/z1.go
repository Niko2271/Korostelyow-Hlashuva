package main

import (
	"fmt"
	"time"
)

func main() {
	pr := make(chan string)
	go sim(pr)
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Превышено время обработки запроса")
	case result := <-pr:
		fmt.Printf("Успешно, ваш результат:", result)
	}
}

func sim(pr chan<- string) {
	time.Sleep(3 * time.Second)
	pr <- "Всем ку"
}
