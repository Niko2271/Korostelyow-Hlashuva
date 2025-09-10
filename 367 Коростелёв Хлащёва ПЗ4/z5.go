package main

import (
 "crypto/md5"
 "fmt"
 "sync"
)

func main() {
 znacheniya := []string{"qqq", "www", "eee", "rrr", "ttt", "yyy"}
 kanal := make(chan bool, 3)
 var krutFgnua sync.WaitGroup
 
 for _, text := range znacheniya {
  krutFgnua.Add(1)
  kanal <- true 
  go func(t string) {
   defer krutFgnua.Done()
   defer func() { <-kanal }() 
   
   hash := md5.Sum([]byte(t))
   fmt.Printf("Хэээш '%s': %x\n", t, hash)
  }(text)
 }
 krutFgnua.Wait()
}
