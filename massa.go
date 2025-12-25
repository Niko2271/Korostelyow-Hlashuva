package main

import (
	"fmt"
)

func main() {
	const USD float64 = 1
	const EUR = 0.8645
	const RUB = 81.23

	EUR_RUB := RUB / EUR
	fmt.Printf("1 евро = %.2f рублей", EUR_RUB)
	EUR_USD := USD / EUR
	RUB_USD := USD / RUB
	RUB_EUR := EUR / RUB

	var a int
	var amount float64

	fmt.Printf("1. USD -> EUR\n2. USD -> RUB\n3. EUR -> USD\n4. EUR -> RUB\n5. RUB -> USD\n6. RUB -> EUR\n")
	fmt.Println("Напишите цифру конвертации: ")
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Введите сколько USD: ")
		fmt.Scan(&amount)
		c := amount * EUR
		fmt.Printf(" %.2f USD = %.2f EUR", amount, c)
	case 2:
		fmt.Println("Введите сколько USD: ")
		fmt.Scan(&amount)
		c := amount * RUB
		fmt.Printf(" %.2f USD = %.2f RUB", amount, c)
	case 3:
		fmt.Println("Введите сколько EUR: ")
		fmt.Scan(&amount)
		c := amount * EUR_USD
		fmt.Printf(" %.2f EUR = %.2f USD", amount, c)
	case 4:
		fmt.Println("Введите сколько EUR: ")
		fmt.Scan(&amount)
		c := amount * EUR_RUB
		fmt.Printf(" %.2f EUR = %.2f RUB", amount, c)
	case 5:
		fmt.Println("Введите сколько RUB: ")
		fmt.Scan(&amount)
		c := amount * RUB_USD
		fmt.Printf(" %.2f RUB = %.2f USD", amount, c)
	case 6:
		fmt.Println("Введите сколько RUB: ")
		fmt.Scan(&amount)
		c := amount * RUB_EUR
		fmt.Printf(" %.2f RUB = %.2f EUR", amount, c)
	default:
		fmt.Println("Ошибко нет такой цифры")
	}
}
