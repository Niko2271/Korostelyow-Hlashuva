package main

import ( "fmt"
	"time" )

func main(){
	tick := time.Tick(time.Millisecond * 200)
	pr := make(chan int, 15)
	for iv := 1; iv <= 15; iv++{
	<- tick		
	fmt.Println("Подсчёт: ", iv)
	pr <- iv
}
}


