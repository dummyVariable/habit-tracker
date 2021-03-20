package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type habitJSONSchema struct {
	Habits  []string `json:"habits"`
	Entries []habit  `json:"entries"`
}

type habitJSONStore struct {
	dataFilename string
}

func newJSONStore() habitJSONStore {
	return habitJSONStore{
		dataFilename: "habit.json",
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
	defer f.Close()

	if err != nil {
		return err
	}

	return nil

}

func readData() habitJSONSchema {

	data, err := ioutil.ReadFile("habit.json")
	checkErr(err)

	var habits habitJSONSchema

	err = json.Unmarshal(data, &habits)
	checkErr(err)

	return habits
}

func writeData(habits habitJSONSchema) {
	jsonData, err := json.Marshal(habits)
	checkErr(err)

	err = ioutil.WriteFile("habit.json", jsonData, 0644)
	checkErr(err)
}
