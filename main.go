package main

import (
	. "fmt"
	. "time"
)
type Task struct {
		ID   int
		Text string
		Done bool
}

func printTask(t Task){
		Sleep(2*Second)
		Print("\n", "ID: ", t.ID)
		Sleep(2*Second)
		Print("\n", "Задание: ", t.Text)
		Sleep(2*Second)
		Print("\n", "Статус: ", t.Done, "\n")
}
func inputTask(id int, name string, status bool) Task{
	task := Task{id, name, status}
	return task
}

func main() {
	var name, text string
	var id int
	var status bool
	Print("Введите ваше имя: ")
	Scan(&name)
	Print("Привет, ", name, "! Это наш менеджер задач")
	Print("\n", Now().Year(), "-", Now().Month(), "-", Now().Day())
	Print("\n", Now().Hour(), ":", Now().Minute(), ":", Now().Second(), "\n")
	Print("Введите информацию о вашей задаче: ")
	Scan(&id, &text, &status)
	tasks := make([]Task, 10)
	tasks[0] = inputTask(id, text, status)
	tasks = append(tasks, Task{0, "Покушать", false})
	printTask(tasks[0])
	// printTask(tasks[1])
	Print(tasks[1])
}
