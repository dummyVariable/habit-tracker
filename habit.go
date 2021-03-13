package main

import (
	"time"
)

//task contains the per completion info of a task for an habit
type task struct {
	name        string
	reps        string
	completedAt time.Time
	bestReps    string
	lastBestAt  time.Time
}

//habit contains the habit name, habit tasks and the habit's adoption rate, and
//habit started.
type habit struct {
	name         string
	adoptionRate int
	startedAt    time.Time
}

//habitDB is implemented by datastore for the habits
type habitDB interface {
	addHabit(habit habit) error
	removeHabit(habitName string) error

	addTask(habitName string, task task) error
	removeTask(habitName, taskName string) error
}
