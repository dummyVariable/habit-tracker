package main

import (
	"time"
)

//habit contains the habit name, and the habit's adoption rate, habit started and current streak.
type habit struct {
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AdoptionRate     int       `json:"adoptionRate"`
	CreatedAt        time.Time `json:"createdAt"`
	Streak           int       `json:"streak"`
	LastCompletionAt time.Time `json:"lastCompletionAt`
}

//habitDB is implemented by datastore for the habits
type habitDB interface {
	addHabit(habit habit) error
	removeHabit(habitName string) error
	completeHabit(habitName string) error
	reportHabit(habitName string) error
}
