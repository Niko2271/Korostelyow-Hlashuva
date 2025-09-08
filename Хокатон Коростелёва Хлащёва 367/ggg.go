package main

import (
    "fmt"
    "math/rand"
    
)

type Order struct {
    chislo int
}

var orders = make(map[int]Order)

func vibor() {
    var chislo int
    fmt.Println("Добрый день, я чат-бот помощник. Выберите цифру, а я вам помогу")
    fmt.Println("Выберите цифру действия:")
    fmt.Println("1. Фильм по жанру;")
    fmt.Println("2. Фильм по настроению;")
    fmt.Println("3. Музыка по жанру;")
    fmt.Println("4. Музыка по настроению;") 
    fmt.Println("5. Рассказать анекдот;")
    fmt.Println("6. Выбор мерч.")
    fmt.Scan(&chislo)

    menu(chislo)
    orders[chislo] = Order{chislo}
}

func menu(chislo int) {
    switch chislo {
    case 1:
        fj()
    case 2:
        fn()
    case 3:
        mj()
    case 4:
        mn()
    case 5:
        anecdot()
    case 6:
        merch()
    default:
        fmt.Println("Неверно >:( ")
    }
}


func fj() {
    var hhru int
    fmt.Println("1. Боевик")
    fmt.Println("2. Комедии")
    fmt.Println("3. Драма")
    fmt.Scan(&hhru)
    switch hhru{
    case 1:
        fmt.Println("Возвращение гремлинов, Лысый нянь, Cтаруха с ножом")
    case 2:
        fmt.Println("Холоп 2,Папины дочки. Новогодние, Маме снова 17")
    case 3:
        fmt.Println("Ворошиловский стрелок, Огонь, Доктор Хаус")
    default:
        fmt.Println("Неверно >:( ")
    }
    
}

func fn() {
    var hhru int
    fmt.Println("1. Весёлое")
    fmt.Println("2. Грустное")
    fmt.Println("3. Хорошее")
    fmt.Scan(&hhru)
    switch hhru{
    case 1:
        fmt.Println("Операция „Панда“,Анора, Герой наших снов")
    case 2:
        fmt.Println("Нелюбовь, Подбросы, История одного назначения")
    case 3:
        fmt.Println("«Я худею», «Всегда говори „да!“», «День сурка»")
    default:
        fmt.Println("Неверно >:( ")
    }
}

func mj() {
    var hhru int
    fmt.Println("1. ХипХоп")
    fmt.Println("2. Рэп")
    fmt.Println("3. Рок")
    fmt.Scan(&hhru)
    switch hhru{
    case 1:
        fmt.Println("Пусть Бог будет рядом с твоими планами,NAME")
    case 2:
        fmt.Println("Fake ID, Наследство")
    case 3:
        fmt.Println("Пачка сигарет,Группа крови")
    default:
        fmt.Println("Неверно >:( ")
    }
}

func mn() {
    var hhru int
    fmt.Println("1. Весёлое")
    fmt.Println("2. Грустное")
    fmt.Println("3. Хорошее")
    fmt.Scan(&hhru)
    switch hhru{
    case 1:
        fmt.Println("Операция „Панда“,Анора, Герой наших снов")
    case 2:
        fmt.Println("Нелюбовь, Подбросы, История одного назначения")
    case 3:
        fmt.Println("«Я худею», «Всегда говори „да!“», «День сурка»")
    default:
        fmt.Println("Неверно >:( ")
    }
}

func anecdot() {
    num := rand.Intn(5)
    switch num{
    case 1:
        fmt.Println("Вместо того, чтобы у себя в Алжире, Марокко и Тунисе строить жизнь, как во Франции, люди приезжают во Францию, чтобы там создать себе такую жизнь, как в Алжире, Марокко и Тунисе. Это странно. Ещё более странно заставлять при этом и французов жить, как в Алжире, Марокко и Тунисе.")
    case 2:
        fmt.Println("Лозунг Задушим коррупцию был признан экстремистским как призывающий к насильственному свержению существующего строя.")
    case 3:
        fmt.Println("Увидев на холодильнике всего два магнитика - из Магадана и Воркуты, воры покормили кота и вымыли посуду.")
    case 4:
        fmt.Println("Когда в стране коррупции нет — микролитражки мчат по хайвеям. Когда коррупция — Бентли тащатся по бездорожью. Всё просто, брат.")
    case 5:
        fmt.Println("Интересно, а если провести обыск у всего руководства ФСБ - можно будет обратно понизить пенсионный возраст?")
    default:
        fmt.Println("Неверно >:( ")
    }
    
}

func merch() {
    fmt.Println("Выбор мерча")
}

func main() {
    vibor()
}