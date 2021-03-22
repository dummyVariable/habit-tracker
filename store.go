package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

const jsonFileName = "habit.json"

type habitJSONSchema struct {
	Habits  map[string]habit `json:"habits"`
	Entries []habit          `json:"entries"`
}

type habitJSONStore struct {
	dataFilename string
}

func newJSONStore() habitJSONStore {
	return habitJSONStore{
		dataFilename: jsonFileName,
	}
}

func isJSONExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func createJSONFile(filename string) error {
	if isJSONExists(filename) {
		return ErrJSONFileAlreadyExists
	}

	f, err := os.Create(filename)
	f.Write([]byte("{}")) //Err raised when unmarshalling empty json file!!
	defer f.Close()

	if err != nil {
		return err
	}

	return nil

}

func readData() habitJSONSchema {

	data, err := ioutil.ReadFile(jsonFileName)
	checkErr(err)

	var habits habitJSONSchema

	err = json.Unmarshal(data, &habits)
	checkErr(err)

	if habits.Habits == nil {
		habits.Habits = make(map[string]habit)
	}

	return habits
}

func writeData(habits habitJSONSchema) {
	jsonData, err := json.Marshal(habits)
	checkErr(err)

	err = ioutil.WriteFile(jsonFileName, jsonData, 0644)
	checkErr(err)
}

func isHabitExists(habitName string) bool {
	habits := readData()

	if _, present := habits.Habits[habitName]; present {
		return true
	}

	return false
}

func (db habitJSONStore) addHabit(newHabit habit) error {

	if !isJSONExists(jsonFileName) {
		return ErrJSONFileNotExists
	}

	if isHabitExists(newHabit.Name) {
		return ErrHabitAlreadyExists
	}

	newHabit.AdoptionRate = 0
	newHabit.CreatedAt = time.Now()

	habits := readData()

	habits.Habits[newHabit.Name] = newHabit
	habits.Entries = append(habits.Entries, newHabit)

	writeData(habits)

	return nil

}

func (db habitJSONStore) removeHabit(habitName string) error {

	if !isJSONExists(jsonFileName) {
		return ErrJSONFileNotExists
	}

	if !isHabitExists(habitName) {
		return ErrHabitNotExists
	}

	habits := readData()

	delete(habits.Habits, habitName)

	writeData(habits)

	return nil
}

func (db habitJSONStore) completeHabit(habitName string) error {

	if !isJSONExists(jsonFileName) {
		return ErrJSONFileNotExists
	}

	if !isHabitExists(habitName) {
		return ErrHabitNotExists
	}

	habits := readData()

	completedHabit := habits.Habits[habitName]
	currentTime := time.Now()

	if completedHabit.LastCompletionAt.Sub(currentTime).Hours() < 25 {
		completedHabit.Streak++
	} else {
		completedHabit.Streak = 0
	}

	completedHabit.LastCompletionAt = currentTime

	habits.Entries = append(habits.Entries, completedHabit)

	writeData(habits)

	return nil
}
