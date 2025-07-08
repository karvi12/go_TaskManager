package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	. "time"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

func printTask(t Task) {
	Print("\n", "ID: ", t.ID)
	Print("\n", "Задание: ", t.Text)
	Print("\n", "Статус: ", t.Done, "\n")
}
func inputTask(id int, name string, status bool) Task {
	return Task{id, name, status}
}

func main() {
	var name string
	var id int
	var status string

	reader := bufio.NewReader(os.Stdin)

	Print("Введите ваше имя: ")
	Scan(&name)
	Print("Привет, ", name, "! Это наш менеджер задач")
	Print("\n", Now().Year(), "-", Now().Month(), "-", Now().Day())
	Print("\n", Now().Hour(), ":", Now().Minute(), ":", Now().Second(), "\n")

	Print("Введите ID задачи: ")
	Scan(&id)

	reader.ReadString('\n')

	Print("Введите описание задачи: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	Print("Введите статус задачи (true/false): ")
	Scan(&status)

	statusBool := status == "true"

	tasks := []Task{}
	tasks = append(tasks, inputTask(id, text, statusBool))
	printTask(tasks[0])
}
