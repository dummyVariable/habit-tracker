package main

import (
	"time"

	"github.com/google/uuid"
)

type (
	habitID uuid.UUID
	taskID  uuid.UUID
)

//habitDBStore implements habitDB in-memory
type habitDBStore struct {
	habits map[habitID]habit
	tasks  map[taskID]task

	habitTaskMap map[habitID][]taskID
}

func newStore() habitDBStore {
	return habitDBStore{
		habits: make(map[habitID]habit),
		tasks:  make(map[taskID]task),

		habitTaskMap: make(map[habitID][]taskID),
	}
}

func (db *habitDBStore) addHabit(newHabit habit) error {
	ID := uuid.New()
	newHabit.ID = ID
	newHabit.adoptionRate = 0
	newHabit.startedAt = time.Now()

	db.habits[habitID(newHabit.ID)] = newHabit
	return nil
}
