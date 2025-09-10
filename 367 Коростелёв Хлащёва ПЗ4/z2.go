package main

import "fmt"

func main(){
 jobs:= make (chan int)
 results:= make(chan int)
 
 go func(){
  for x :=1; x<=10; x++ {
   jobs <- x
  }
 close(jobs)
 }()

 go func(){
  for x := range jobs{
   results <- x*x
  }
 close(results)
 }()

 for x := range results{
  fmt.Println(x)
 }
}