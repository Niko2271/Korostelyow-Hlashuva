package main

import (
	"fmt"
)

type Task struct {
	Name   string
	Vipoln bool
}
type TaskManager struct {
	tasks []Task
}

func (t *TaskManager) Status(name string, done bool) {
	for i := range t.tasks {
		if t.tasks[i].Name == name {
			t.tasks[i].Vipoln = done
			status := "невыполн"
			if done {
				status = "выполн"
			}
			fmt.Printf("Задача '%s' отмечена как %s\n", name, status)
			return
		}
	}
	fmt.Printf("Задача '%s' не найдена\n", name)
}

func (t *TaskManager) Add(name string) {
	t.tasks = append(t.tasks, Task{Name: name})
	fmt.Printf("Задача добавлена")
}

func (t *TaskManager) Delete(name string) {
	for i := 0; i < len(t.tasks); i++ {
		if t.tasks[i].Name == name {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			fmt.Println("Удалили")
			return
		}
	}
	fmt.Println("Нет такого")
}

func (t *TaskManager) Spisok() {
	fmt.Println("\nСписок задач: ")
	for i, task := range t.tasks {
		status := " "
		if task.Vipoln {
			status = "ес"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Name)
	}
}

func main() {
	t := TaskManager{}

	t.Add("купить молоко")
	t.Add("сделать сп")
	t.Spisok()
	t.Delete("сделать сп")
	t.Status("купить молоко", true)
	t.Spisok()
}
