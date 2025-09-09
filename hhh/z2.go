package main

import (
	"fmt"
	"time"
	"sync"
)

func countr(){
	for c := 0; c < 10; c++{
		out <- c
	}
	close(out)
}


func main() {
	j := make(chan int)
	res := make(chan int)

	go countr(j)
	go sqr(sqr, j)
	priter(sqr)
	там очень сильно пахло электронками
}