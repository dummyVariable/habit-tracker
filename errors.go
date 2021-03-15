package main

import "errors"

var (
	ErrHabitNotExists     = errors.New("Habit not exists")
	ErrHabitAlreadyExists = errors.New("Habit already exists")
	ErrTaskNotExists      = errors.New("Task not exists")
	ErrTaskAlreadyExists  = errors.New("Task already exists")
)
