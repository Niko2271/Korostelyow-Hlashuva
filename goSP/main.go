package main

import (
	"fmt"
	"math/rand"
)

func menu() int {
	var t int
	fmt.Println("Привет! Я чат-бот помощник, выбери цифру того, что тебя интересует.")
	fmt.Println("1. Рекомендация фильма по жанру ")
	fmt.Println("2. Рекомендация фильма по настроению ")
	fmt.Println("3. Рекомендация музыки по жанру ")
	fmt.Println("4. Рекомендация музыки по настроению ")
	fmt.Println("5. Рассказать анекдот ")
	fmt.Println("6. Каталог мерча ")
	fmt.Scan(&t)

	return t
}

func vibor(t int) {
	if t == 1 {
		fj()
	}
	if t == 2 {
		fn()
	}
	if t == 3 {
		mj()
	}
	if t == 4 {
		mn()
	}
	if t == 5 {
		ane()
	}
	if t == 6 {
		merch()
	}
}

func fj() {
	var o int
	fmt.Println("Супер! Выбери интересующий тебя жанр:")
	fmt.Println("1. Комедия")
	fmt.Println("2. Ужасы")
	fmt.Println("3. Боевик")
	fmt.Println("4. Фантастика")
	fmt.Scan(&o)

	if o == 1 {
		fmt.Println("Отличный выбор! Вот несколько фильмов в данном жанре:")
		fmt.Printf("«Жмурки» (2005)\n«Новогодний шеф» (2023)\n«Отель „Гранд Будапешт“» (2014) ")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if o == 2 {
		fmt.Println("Отличный выбор! Вот несколько фильмов в данном жанре:")
		fmt.Printf("«Заклятье. Шёпот ведьм» (2024)\n«Крик. Сезон Призраков» (2024)\n«Все мои друзья мертвы» (2024)")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if o == 3 {
		fmt.Println("Отличный выбор! Вот несколько фильмов в данном жанре:")
		fmt.Printf("«Форсаж» (2001)\n«Крепкий орешек» (1988)\n«Проект „Золотой глаз“»")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if o == 4 {
		fmt.Println("Отличный выбор! Вот несколько фильмов в данном жанре:")
		fmt.Printf("«Хищник: Миссия „Осирис“» (2025)\n«Жизнь Чака» (2025)\n«М3ГАН 2.0» (2025)")
		fmt.Printf("\nПриятного просмотра :)")
	}
}
func fn() {
	var k int
	fmt.Println("Супер! Какое у тебя настроение сейчас?")
	fmt.Println("1. Грустное")
	fmt.Println("2. Радостное")
	fmt.Println("3. Нейтральное")
	fmt.Println("4. Рабочее")
	fmt.Scan(&k)

	if k == 1 {
		fmt.Println("Не переживай всё буднт хорошо! Подобрал несколько фильмов под твоё настроение:")
		fmt.Printf("«Один день» (2011)\n«Время между нами» (2022)\n«Я найду тебя» (2019)")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if k == 2 {
		fmt.Println("Это прекрасно! Подобрал несколько фильмов под твоё настроение:")
		fmt.Printf("«Приключения Паддингтона» (2014)\n«Каникулы строгого режима» (2009)\n«Вторая жизнь Уве» (2015)")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if k == 3 {
		fmt.Println("Хороший выбор! Подобрал несколько фильмов под твоё настроение:")
		fmt.Printf("«Красотка на всю голову» (2018)\n«2 дня» (2011)\n«Итальянские каникулы» (2020)")
		fmt.Printf("\nПриятного просмотра :)")
	}
	if k == 4 {
		fmt.Println("Подобрал несколько фильмов для фона:")
		fmt.Printf("«Стажер» (2015)\n«Простые сложности» (2009)\n«Повар на колёсах»")
		fmt.Printf("\nПриятного просмотра :)")
	}
}
func mj() {
	var x int
	fmt.Println("Супер! Выбери интересующий тебя жанр музыки:")
	fmt.Println("1. Поп")
	fmt.Println("2. Рэп и Хип-хоп")
	fmt.Println("3. Альтернатива")
	fmt.Println("4. Рок")
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("ДЫМ GRILLYAZH\nКрасками HOLLYFLAME\nThe One That You Love LP")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if x == 2 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("ТАЯЛИ f0lk\nКЛАССИКА Цинк Уродов, H8.HOOD\n50 на 50 pyrokinesis")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if x == 3 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("Peoples' Republic of Desire The Vaccines\nя.и.п.н.х.м. джинсы тарковского\nПоцелуи Neverlove")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if x == 4 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("Хейла-Боппа Unnamed Feeling\nКонец Дайте танк (!)\nУлыбайся Norma Tale")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
}
func mn() {
	var r int
	fmt.Println("Супер! Выбери какое у тебя сейчас настроение:")
	fmt.Println("1. Радостное")
	fmt.Println("2. Спокойное")
	fmt.Println("3. Грустное")
	fmt.Println("4. Мечтатльное")
	fmt.Scan(&r)

	if r == 1 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("Up Against Me LP\n KULTURA Апология\nGEMINI BLK ODYSSY, Jackie Giroux")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if r == 2 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("How Onative, Zubi\nUntil I Found You Stephen Sanchez\nThe Bruise Damiano David, Suki Waterhouse")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if r == 3 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("СИНДРОМ f0lk\nПолчаса HOLLYFLAME\nPrayers KADI, Miyagi")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
	if r == 4 {
		fmt.Println("Отличный выбор! Несколько песен для тебя:")
		fmt.Printf("Love BAGARDI\nНезаметно Три дня дождя, LYRIQ\nБудь со мной Colorit")
		fmt.Printf("\nПриятного прослушивания >.< ")
	}
}
func ane() {
	num := rand.Intn(4)
	if num == 1 {
		fmt.Println("Штирлиц всегда спал как убитый. Его даже пару раз обводили мелом")
	}
	if num == 2 {
		fmt.Printf("Штирлиц выкрал у Бормана важные бумаги и пошёл фотографировать их в туалет. Вдруг в унитазе возникает Борман:\n— Штирлиц, вы окружены. Со мной тысяча солдат. \n— А со мной техника, — сказал Штирлиц, дёргая верёвку")
	}
	if num == 3 {
		fmt.Println("Штирлиц в третий раз терял передатчик, и его весь вечер бил озноб. Озноб был радистом, и без передатчика ему было никак")
	}
	if num == 4 {
		fmt.Println("Штирлиц шёл по улице, когда внезапно прямо перед ним что-то упало. Штирлиц поднял глаза. Это были глаза профессора Плейшнера")
	}

}
func merch() {
	var h int
	fmt.Println("Каталог мерча:")
	fmt.Println("1. Кофты")
	fmt.Println("2. Тетради")
	fmt.Println("3. Игрушки")
	fmt.Scan(&h)

	if h == 1 {
		var e string
		fmt.Println("Какой цвет предпочитаете?")
		fmt.Scan(&e)
		var n string
		fmt.Println("Введите свой размер(цифрами)")
		fmt.Scan(&n)
		var y string
		fmt.Println("Для кого(тип одежды: м, ж, унисекс)")
		fmt.Scan(&y)

		var f string
		fmt.Println("Ваш заказ: кофта цвет:", e, "размер:", n, "пол:", y, ". Оформляем заказ?")
		fmt.Scan(&f)
		if f == "да" {
			fmt.Println("Заказ оформлен! Скоро с вами свяжется оператор.")
		}
		if f == "нет" {
			fmt.Println("Печально( Хорошего дня!")
		}
	}

	if h == 2 {
		var e string
		fmt.Println("Какой цвет предпочитаете?")
		fmt.Scan(&e)
		var n string
		fmt.Println("Введите количество листов")
		fmt.Scan(&n)
		var y string
		fmt.Println("Введите коллекцицию: ")
		fmt.Scan(&y)
		var g int
		fmt.Println("Сколько тетрадей нужно?")
		fmt.Scan(&g)

		var f string
		fmt.Println("Ваш заказ: тетради цвет:", e, "листов:", n, "коллекция", y, "количество", g, ". Оформляем заказ?")
		fmt.Scan(&f)
		if f == "да" {
			fmt.Println("Заказ оформлен! Скоро с вами свяжется оператор.")
		}
		if f == "нет" {
			fmt.Println("Печально( Хорошего дня!")
		}
	}

	if h == 3 {
		var e string
		fmt.Println("Введите коллекцию: ")
		fmt.Scan(&e)
		var n string
		fmt.Println("Введите персонажа: ")
		fmt.Scan(&n)
		var y string
		fmt.Println("Введите размер(20см, 25см, 30см):")
		fmt.Scan(&y)

		var f string
		fmt.Println("Ваш заказ: ишрушка коллекция:", e, "персонаж:", n, "размер", y, ". Оформляем заказ?")
		fmt.Scan(&f)
		if f == "да" {
			fmt.Println("Заказ оформлен! Скоро с вами свяжется оператор.")
		}
		if f == "нет" {
			fmt.Println("Печально( Хорошего дня!")
		}
	}
}

func main() {
	u := menu()
	vibor(u)
}
