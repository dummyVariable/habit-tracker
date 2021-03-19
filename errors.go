package main

import "errors"

var (
	ErrHabitNotExists     = errors.New("Habit not exists")
	ErrHabitAlreadyExists = errors.New("Habit already exists")
)
