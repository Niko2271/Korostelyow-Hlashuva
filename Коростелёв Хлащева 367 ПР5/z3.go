package main

import (
	"fmt"
	"time"
)

func main() {
	pr := make(chan string)

	go sim(pr)

	select {
	case result := <-pr:
		fmt.Println("Успешно, ваш результат:", result)
	default:
		fmt.Println("файл пуст")
	}
	time.Sleep(3 * time.Second)
	select {
	case result := <-pr:
		fmt.Println("Успешно, ваш результат:", result)
	default:
		fmt.Println("файл пуст")
	}
}

func sim(pr chan<- string) {
	time.Sleep(2 * time.Second)
	pr <- "Всем ку"
}
