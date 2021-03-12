package main

import "time"

//task contains the per completion info of a task for an habit
type task struct {
	name        string
	reps        string
	completedAt time.Time
}
