package main

import (
	"fmt"
	"sync"
	"time"
)

type QQQ struct {
	sync.Mutex
	iti []string
}

func (q *QQQ) Add(s string) {
	q.Lock()
	defer q.Unlock()
	q.iti = append(q.iti, s)
}

func (q *QQQ) Remove() string {
	q.Lock()
	defer q.Unlock()
	if len(q.iti) == 0 {
		return ""
	}
	item := q.iti[0]
	q.iti = q.iti[1:]
	return item
}

func main() {
	q := &QQQ{}
	
	go func() {
		for i := 0; i < 4; i++ {
			q.Add(fmt.Sprintf("задачу %d", i))
			fmt.Printf("Добавил задачу: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Println("Получил:", q.Remove())
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Завершение")
}
