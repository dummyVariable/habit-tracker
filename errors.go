package main

import "errors"

var (
	ErrHabitNotExists        = errors.New("Habit not exists")
	ErrHabitAlreadyExists    = errors.New("Habit already exists")
	ErrJSONFileAlreadyExists = errors.New("JSON file already exists")
)
