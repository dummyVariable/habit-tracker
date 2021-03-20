package main

type habitJSONStore struct {
	dataFilename string
}

func newJSONStore() habitJSONStore {
	return habitJSONStore{
		dataFilename: "habit.json",
	}
}
