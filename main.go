package main

import (
	"fmt"
)

var numTasks int

type Tasks struct {
	taskName     string
	numPomodoros int
}

func main() {

	fmt.Println("------------------Black Pomodoro------------------")
	fmt.Println("---------------------Welcome!---------------------")
	fmt.Println("------------Let's fill today's session------------")

	getData()

}

func getData() {

	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	fmt.Print("How many tasks for this session? -> ")
	fmt.Scan(&numTasks)

	totalTasks := make([]Tasks, 0)
	totalPomodoros := make([]int, 0)

	task := ""
	pomodoroNum := 0
	pomodoroLen := 0
	currentTask := Tasks{
		taskName:     "",
		numPomodoros: 0,
	}

	for i := 0; i < numTasks; i++ {
		fmt.Printf("Name of task #%v? -> ", i+1)
		fmt.Scan(&task)
		fmt.Printf("How many Pomodoros for %v task -> ", task)
		fmt.Scan(&pomodoroNum)
		fmt.Printf("Amount of time for each %v Pomodoros? -> ", task)
		fmt.Scan(&pomodoroLen)

		currentTask.taskName = task
		currentTask.numPomodoros = pomodoroNum

		totalPomodoros = append(totalPomodoros, pomodoroLen)
		totalTasks = append(totalTasks, currentTask)

		fmt.Println(totalTasks)
		fmt.Println(totalPomodoros)
	}
}

// func pomodoros() {
// 	return
// }

// func checkForBreaks() {
// 	return
// }

// func applyBreak() {
// 	return
// }
