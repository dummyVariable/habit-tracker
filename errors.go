package main

import (
	"errors"
	"log"
)

var (
	ErrHabitNotExists        = errors.New("Habit not exists")
	ErrHabitAlreadyExists    = errors.New("Habit already exists")
	ErrJSONFileAlreadyExists = errors.New("JSON file already exists")
	ErrJSONFileNotExists     = errors.New("Json file not exists")
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
