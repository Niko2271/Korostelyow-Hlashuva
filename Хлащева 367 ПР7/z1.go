package main

import (
  "fmt"
  "errors"
)

type BankAccount struct{
    AccountNumber int
    holderName string
    balance float64
}


func (BA * BankAccount) Deposit(popolnenie float64){
  fmt.Println("Операция пополнения")
  if popolnenie > 0 {
    BA.balance += popolnenie
  }
}

func (BA *BankAccount) Withdraw(spisanie float64) error {
  fmt.Println("Операция списания")
 if spisanie <= 0 {
  return errors.New("Вы не можете ввести отрицательное число")
 }
 if spisanie > BA.balance {
  return errors.New("Вашего баланса не достаточно для выполнения операции")
 }
 BA.balance -= spisanie
 return nil
}

func (BA *BankAccount) GetBalance() float64 {
 return BA.balance
}

func (BA *BankAccount) PrintBalance() {
 fmt.Printf("Ваш счёт составляет: %.2f\n", BA.balance)
}

func main(){
  account := BankAccount{
    AccountNumber: 14324,
    holderName: "Деньгин Роман Михайлович",
    balance:    1200.0,
  }
  account.Deposit(679.59)
  account.PrintBalance()

  err := account.Withdraw(300.0)
   if err != nil {
    fmt.Println("Ошибка:", err)
   } else {
  account.PrintBalance()
 }


   err = account.Withdraw(3333.33)
   if err != nil {
    fmt.Println("Ошибка:", err)
   } else {
    account.PrintBalance()
  }
}
