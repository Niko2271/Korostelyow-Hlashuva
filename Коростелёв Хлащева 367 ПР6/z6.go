package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex

func logSafe(message string) {
	l.Lock()
	defer l.Unlock()
	fmt.Println(message)
}

func main() {
	var w sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		w.Add(1)
		go func(id int) {
			defer w.Done()
			for j := 0; j < 2; j++ {
				logSafe(fmt.Sprintf("Горутина %d: лог %d", id, j))
			}
		}(i)
	}
	
	w.Wait()
}
