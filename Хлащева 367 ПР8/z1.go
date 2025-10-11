package main

import  "fmt"

type Logger interface{
  Info(msg string)
  Error(msg string)
  Debug(msg string)
}

type FileLogger struct{}

type ConsoleLogger struct{}

func(ConsoleLogger) Info(msg string){
  fmt.Println("info", msg)
}
func(ConsoleLogger) Error(msg string){
  fmt.Println("error", msg)
}
func(ConsoleLogger) Debug(msg string){
  fmt.Println("debug", msg)
}

func(FileLogger) Info(msg string){
  fmt.Println("file info", msg)
}
func(FileLogger) Error(msg string){
  fmt.Println("file error", msg)
}
func(FileLogger) Debug(msg string){
  fmt.Println("file debug", msg)
}

func main(){
  var log Logger = ConsoleLogger{}
  log.Info("тест")
  log.Error("ошибочка")
  log.Debug("отладочка")
}
