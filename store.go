package main

import (
	"encoding/json"
	"fmt"
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

func calcAdoptionRate(habitName string, completedHabit habit) int {

	var streak, missed float32

	habits := readData()

	habitEntries := make([]habit, 0)

	for _, habit := range habits.Entries {
		if habitName == habit.Name {
			habitEntries = append(habitEntries, habit)
		}
	}

	habitEntries = append(habitEntries, completedHabit)

	var prev, temp time.Time

	for i, entry := range habitEntries {

		if i == 0 {
			temp = habits.Habits[habitName].CreatedAt
		} else {
			temp = prev
		}

		if entry.LastCompletionAt.Sub(temp).Hours() < 25 {
			streak++
		} else {
			missed++
		}

		prev = entry.LastCompletionAt
	}

	if missed == 0 {
		return 100
	}

	return int((streak / (streak + missed)) * 100)

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

	if currentTime.Sub(completedHabit.LastCompletionAt).Hours() < 25 {
		completedHabit.Streak++
	} else {
		completedHabit.Streak = 0
	}

	completedHabit.LastCompletionAt = currentTime
	completedHabit.AdoptionRate = calcAdoptionRate(habitName, completedHabit)

	habits.Habits[completedHabit.Name] = completedHabit
	habits.Entries = append(habits.Entries, completedHabit)

	writeData(habits)

	return nil
}

func (db habitJSONStore) reportHabit(habitName string) error {

	if !isJSONExists(jsonFileName) {
		return ErrJSONFileNotExists
	}

	if !isHabitExists(habitName) {
		return ErrHabitNotExists
	}

	habits := readData()
	currentHabit := habits.Habits[habitName]

	fmt.Printf("\nHabit : %v\n", currentHabit.Name)

	if currentHabit.Description != "" {
		fmt.Printf("Description : %v\n", currentHabit.Description)
	}

	fmt.Printf("\nAdoption Rate : %v\n", currentHabit.AdoptionRate)
	fmt.Printf("Created At : %v\n", currentHabit.CreatedAt.String())
	fmt.Printf("Current Streak : %v\n", currentHabit.Streak)

	if currentHabit.LastCompletionAt.IsZero() {
		fmt.Println("Not Started yet")
	} else {
		fmt.Printf("Last Completion at : %v\n", currentHabit.LastCompletionAt.String())
	}

	return nil
}
