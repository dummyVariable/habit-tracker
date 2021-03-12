package main

import (
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
