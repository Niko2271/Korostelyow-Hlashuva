package main

import (
	"fmt"
	"time"
)

func main() {
	sroch := make(chan string, 3)
	obch := make(chan string, 3)

	go Obrabotka(sroch, obch)

	go func() {
		sroch <- "1 Срочный"
		obch <- "2 Обычный"
		sroch <- "3 Срочный"
	}()

	time.Sleep(15 * time.Second)
}

func Obrabotka(sroch chan string, obch chan string) {
	for {
		select {
		case zak := <-sroch:
			ob(zak, true)
			continue
		default:
		}
		select {
		case zak := <-sroch:
			ob(zak, true)
		case zak := <-obch:
			ob(zak, false)
		}
	}
}
func ob(zak string, s bool) {
	if s == true {
		fmt.Println("Заказ", zak, " ---> Обрабатываю задачу")
		time.Sleep(2 * time.Second)
		fmt.Println("Заказ обработан :)")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Println("Заказ", zak, "---> Обрабатываю задачу")
		time.Sleep(5 * time.Second)
		fmt.Println("Заказ обработан :)")
		time.Sleep(1 * time.Second)
	}
}
