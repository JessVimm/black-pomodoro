package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var numTasks int

type Tasks struct {
	taskName     string
	numPomodoros int
}

var totalTasks = make([]Tasks, 0)
var totalPomodoros = make([]int, 0)

func main() {
	fmt.Println("------------------Black Pomodoro------------------")
	fmt.Println("---------------------Welcome!---------------------")
	fmt.Println("------------Let's fill today's session------------")

	getData()

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

func getData() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	fmt.Print("How many tasks for this session? -> ")
	fmt.Scan(&numTasks)

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
}

func pomodoros(pomodoros []int, tasks []Tasks) {
	thisPTaskMinutes := 0
	thisPTaskDone := 0
	totalPomodoros := 0
	pCount := 0

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
				fmt.Printf("Beginning with pomodoro #%v from task %s (%v done of %v total pomodoros)", thisPTaskDone+1, tasks[index].taskName, pCount+1, totalPomodoros)
				fmt.Println()
				countdown(thisPTaskMinutes)
				// Update pomodoros done for this task
				thisPTaskDone++
				// Update total pomodoros done
				pCount++

				// Check for breaks
				checkForBreaks(thisPTaskDone, tasks[index].numPomodoros, pCount, totalPomodoros, thisPTaskMinutes)
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
	fmt.Println("Starting countdown")
	startTime := time.Now()

	finishTime := startTime.Add(time.Minute * time.Duration(minutes))
	currTime := startTime

	for i := 0; currTime.Before(finishTime); i++ {
		currTime = time.Now()
	}
	ring()
}

func checkForBreaks(pTaskDone int, taskPomodoros int, pCount int, totalP int, pMinutes int) {
	if pCount == totalP { // Have we finished the entire session?
		fmt.Println("Done with all pomodoros...") // No need for breaks
	} else if pTaskDone == taskPomodoros { // Have we finished this task pomodoros?
		applyLongBreak(pMinutes) // Have we gone through half of this taks pomodoros?
	} else if pTaskDone == (taskPomodoros / 2) {
		applyLongBreak(pMinutes)
	} else {
		applyRegularBreak(pMinutes)
	}
}

func applyRegularBreak(pMinutes int) {
	fmt.Println("Regular Break Time!")
	// Calculate break time
	var breakTime float32 = 0.0
	breakTime = float32(pMinutes) * 0.20
	// Apply break
	countdown(int(breakTime))
}

func applyLongBreak(pMinutes int) {
	fmt.Println("Long Break Time!")
	// Calculate break time
	var breakTime float32 = 0.0
	if pMinutes > 25 {
		breakTime = float32(pMinutes) * 0.8
	} else if pMinutes <= 25 {
		breakTime = float32(pMinutes) * 0.6
	}
	// Apply break
	countdown(int(breakTime))
}

func ring() {
	// Open
	sound, err := os.Open("Bomberman.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Decode
	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize speaker
	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// Play
	stopEntry := ""
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {})))

	// Stop sound
	fmt.Print("Continue? -> ")
	fmt.Scanln(&stopEntry)

	// Clear
	speaker.Clear()
}
