package main

import "github.com/google/uuid"

type (
	habitID uuid.UUID
	taskID  uuid.UUID
)

type habitDBStore struct {
	habits map[habitID]habit
	tasks  map[taskID]task

	habitTaskMap map[habitID][]taskID
}
