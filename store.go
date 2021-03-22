package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

const jsonFileName = "habit.json"

type habitJSONSchema struct {
	Habits  []string `json:"habits"`
	Entries []habit  `json:"entries"`
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

	for _, habit := range habits.Habits {
		if habit == habitName {
			return true
		}
	}

	return false
}

func (db habitJSONStore) addHabit(habit habit) error {

	if !isJSONExists(jsonFileName) {
		return ErrJSONFileNotExists
	}

	if isHabitExists(habit.Name) {
		return ErrHabitAlreadyExists
	}

	habit.AdoptionRate = 0
	habit.CreatedAt = time.Now()

	habits := readData()
	habits.Habits = append(habits.Habits, habit.Name)
	habits.Entries = append(habits.Entries, habit)
	writeData(habits)

	return nil

}
