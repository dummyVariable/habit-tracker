package main

import (
	"time"
)

//habit contains the habit name, and the habit's adoption rate, habit started and current streak.
type habit struct {
	name             string
	description      string
	adoptionRate     int
	createdAt        time.Time
	streak           int
	lastCompletionAt time.Time
}

//habitDB is implemented by datastore for the habits
type habitDB interface {
	addHabit(habit habit) error
	removeHabit(habitName string) error
	completeHabit(habitName string) error
	reportHabit(habitName string) error
}
