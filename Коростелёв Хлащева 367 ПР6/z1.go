package main

import (
"fmt"    
"time" 
"sync")

var (mu sync.Mutex 
balance int)

func zahod() {
mu.Lock()
balance++
mu.Unlock()
}

func prosmotr() int {
mu.Lock()
b := balance
fmt.Println(b)
mu.Unlock()
return b}

func main(){  
	go zahod()   
	go zahod()    
	go zahod()    
	go zahod()    
	time.Sleep(200 * time.Millisecond)    
	go prosmotr()    
	time.Sleep(200 * time.Millisecond)
}
