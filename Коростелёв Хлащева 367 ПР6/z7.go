package main
import "fmt"

var tovar = map[string]int{}

func dobavlenie_tovara(name string, kolich int) {
    tovar[name] += kolich
}

func prodazha_tovara(name string, kolich int) string {
    if tovar[name] < kolich {
        return "Товара под названием " + name + " недостаточно, Извините за неудобства"
    }
    tovar[name] -= kolich
    return "Товар " + name + " продан"
}

func main() {
    dobavlenie_tovara("Майонез", 7)
    dobavlenie_tovara("Колбаса", 8)
 dobavlenie_tovara("Газировка", 1)
    dobavlenie_tovara("Корм", 4)
 dobavlenie_tovara("Сок", 2)

    
    fmt.Println(prodazha_tovara("Корм", 3))
    fmt.Println(prodazha_tovara("Газировка", 6))
 fmt.Println(prodazha_tovara("Колбаса", 3))
    fmt.Println(prodazha_tovara("Сок", 2))
 fmt.Println(prodazha_tovara("Майонез", 5))
}