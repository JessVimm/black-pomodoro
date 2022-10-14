package main

import (
	"fmt"
	"os"
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
	}

	userReady := ""

	fmt.Print("Ready? (Y/n) -> ")
	fmt.Scan(&userReady)
	if userReady == "Y" || userReady == "y" || userReady == "Yes" || userReady == "yes" {
		pomodoros(totalPomodoros, totalTasks)
	} else {
		fmt.Println("All right! See you next time :)")
		os.Exit(1)
	}
	fmt.Println("All done! You did a great job :D")
	fmt.Println("See you soon!")

}

func pomodoros(pomodoros []int, tasks []Tasks) {
	thisPTaskMinutes := 0
	done := make(chan bool) // Channel
	thisPTaskDone := 0
	totalPomodoros := 0
	currPCount := 0

	// Calculate total Pomodoros for this session
	for n := 0; n < len(pomodoros); n++ {
		totalPomodoros += tasks[n].numPomodoros
	}

	// Get into each Pomodoro
	if len(tasks) > 0 {
		for index := range tasks {
			thisPTaskMinutes = pomodoros[index] // Each pomodoro time for current task
			for pAmount := tasks[index].numPomodoros; pAmount > 0; pAmount-- {
				// Begin countdown
				fmt.Printf("Beginning with pomodoro # %v from task %s", pAmount+1, tasks[index].taskName)
				fmt.Println()
				countdown(thisPTaskMinutes)
				done <- true
				<-done
				// Update pomodoros done for this task
				thisPTaskDone++
				// Update total pomodoros done
				currPCount++

				// Check for breaks
				checkForBreaks(thisPTaskDone, tasks[index].numPomodoros, currPCount, totalPomodoros, thisPTaskMinutes)
			}
			thisPTaskDone = 0
		}
	} else {
		fmt.Println("Ups o.o seems like you didn't enter any tasks for today. Please, try again :)")
		os.Exit(1)
	}
}

func countdown(minutes int) {
	// Start countdown
	// ring
}

func checkForBreaks(pTaskDone int, taskPomodoros int, currPCount int, totalP int, pMinutes int) {
	// Have we finished the entire session? currPCount == totalP Yes? Then...
	// No need for breaks, say goodbye!

	// Have we finished this task pomodoros? pTaskDone == taskPomodoros Yes? then...
	// Apply this task long break
	// applyLongBreak(pMinutes int)

	// Have we gone through half of this taks pomodoros? pTaskDone == (taskPomodoros/2) Yes? then...
	// Apply long break
	// applyLongBreak(pMinutes int)

	// Else, regular break for this task
	// applyRegularBreak(pMinutes int)

}

// func applyRegularBreak(pMinutes int) {
// Calculate break time: breakTime = pMinutes * 0.20
// Apply break countdown(breakTime)
// ring
// }

// func applyLongBreak(pMinutes int) {
// Calculate break time: if pMinutes > 25 then breakTime = pMinutes*0.8, else if pMinutes <= 25 then breakTime = pMinutes*0.6
// Apply break countdown(breakTime)
// ring
// }

// func ring() {
// Do the ring
// Wait for user
// }
