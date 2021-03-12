package main

import (
	"time"
)

type (
	habitName string
	taskName  string
)

//habitDBStore implements habitDB in-memory
type habitDBStore struct {
	habits map[habitName]habit
	tasks  map[taskName]task

	habitTaskMap map[habitName][]taskName
}

func newStore() habitDBStore {
	return habitDBStore{
		habits: make(map[habitName]habit),
		tasks:  make(map[taskName]task),

		habitTaskMap: make(map[habitName][]taskName),
	}
}

func (db *habitDBStore) addHabit(newHabit habit) error {

	newHabit.adoptionRate = 0
	newHabit.startedAt = time.Now()

	db.habits[habitName(newHabit.name)] = newHabit
	return nil
}
