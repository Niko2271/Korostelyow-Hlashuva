package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	HP       int
	Strength int
}

type Enemy struct {
	Name     string
	HP       int
	Strength int
}

func (e *Enemy) Hit() int {
	return rand.Intn(3) + 1
}

func (e *Enemy) Block() int {
	return rand.Intn(3) + 1
}

func (p *Player) Hit() int {
	var hit int
	fmt.Println("Куда бьешь? 1-руки, 2-ноги, 3-голова")
	fmt.Scan(&hit)
	if hit < 1 || hit > 3 {
		hit = 1
	}
	return hit
}

func (p *Player) Block() int {
	var block int
	fmt.Println("Что защищаешь? 1-крылья(руки), 2-ноги, 3-корпус")
	fmt.Scan(&block)
	if block < 1 || block > 3 {
		block = 1
	}
	return block
}

func Fight(p *Player, e *Enemy) {
	round := 1

	for p.HP > 0 && e.HP > 0 {
		//fmt.Printf("\nРаунд %d\n", round)

		playerBlock := p.Block()
		playerHit := p.Hit()
		enemyHit := e.Hit()
		enemyBlock := e.Block()

		fmt.Printf("Ты бьешь в %d, защищаешь %d\n", playerHit, playerBlock)
		fmt.Printf("%s бьет в %d, защищает %d\n", e.Name, enemyHit, enemyBlock)

		if playerHit != enemyBlock {
			e.HP -= p.Strength
			fmt.Printf("Ты попал! У %s осталось %d HP\n", e.Name, e.HP)
		} else {
			fmt.Printf("%s заблокировал удар щитом!\n", e.Name)
		}

		if e.HP > 0 && enemyHit != playerBlock {
			p.HP -= e.Strength
			fmt.Printf("%s попал! У тебя осталось %d HP\n", e.Name, p.HP)
		} else if e.HP > 0 {
			fmt.Println("Ты заблокировал удар крепкой чешуёй")
		}

		round++
	}

	if p.HP > 0 {
		fmt.Printf("\nПобедитель: %s\n", p.Name)
		fmt.Printf("Больше никто не тревожил девушку и город город жил спокойно с потомками мага.")
	} else {
		fmt.Printf("\nПобедитель: %s\n", e.Name)
		fmt.Printf("Город сгорел и погас в ужасе. Девушка исчезла а мэр вскоре умер от ожогов.")
	}
}

func main() {
	fmt.Printf("Предисловие: Однажды в древние века, в темном лесу на окраине города, простудой приболев слегка\n")
	fmt.Printf("казалось из за пустяка колдун искал ученика. Он долго по миру бродтл, но всё таки нашёл его\n")
	fmt.Printf("но парень магу был не мил. Вскоре магу стало ясно, не подходит тот ему и тратить время ни к чему.\n")
	fmt.Printf("Он так расстроился что взял в глухой деревне ученицу, и в драекона привратилш того ученика тупицу.\n")
	fmt.Printf("Летели годы и столетья. Дети те уже не дети. Маг тот умер мы подметим. Свою силу он дал ученице.\n")
	fmt.Printf("Она любила всех людей и многие из горожан ходилиза советом к ней. Но вскоре зависть овладела сердцем\n")
	fmt.Printf("молодого мэра и тот пытался её имя опорочить. Мэр решил поджечь её дом, услышав пламя море вмиг полило \n")
	fmt.Printf("через рамку гобелена. Посмотрев в окно она лишь хмыкнула, прошептав: Ну что же братец выходи на сцену. \n")
	fmt.Printf("Вы играете за дракона.\n")
	fmt.Printf(" \n")

	rand.Seed(time.Now().UnixNano())
	p := &Player{
		Name:     "Дракон",
		HP:       120,
		Strength: 20}
	e := &Enemy{
		Name:     "Мэр",
		HP:       95,
		Strength: 15}

	Fight(p, e)
}
