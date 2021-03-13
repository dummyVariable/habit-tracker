package main

import (
	"errors"
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

func contains(tasks []taskName, key taskName) bool {
	for _, task := range tasks {
		if key == task {
			return true
		}
	}
	return false
}

func (db *habitDBStore) removeHabit(habit string) error {

	delete(db.habits, habitName(habit))
	tasksOfHabit := db.habitTaskMap[habitName(habit)]

	for key := range db.tasks {
		if contains(tasksOfHabit, key) {
			delete(db.tasks, key)
		}
	}
	delete(db.habitTaskMap, habitName(habit))
	return nil

}

func (db *habitDBStore) addTask(habit string, newTask task) error {

	if _, present := db.habits[habitName(habit)]; !present {
		return errors.New("Habit not exists")
	}

	db.tasks[taskName(newTask.name)] = newTask
	db.habitTaskMap[habitName(habit)] = append(db.habitTaskMap[habitName(habit)], taskName(newTask.name))

	return nil
}
func (db *habitDBStore) removeTask(habit, task string) error {

	var index int

	if !contains(db.habitTaskMap[habitName(habit)], taskName(task)) {
		return errors.New("Task not exists")
	}

	for i, key := range db.habitTaskMap[habitName(habit)] {
		if key == taskName(task) {
			index = i
			break
		}

	}
	db.habitTaskMap[habitName(habit)] = append(db.habitTaskMap[habitName(habit)][:index], db.habitTaskMap[habitName(habit)][index+1:]...)
	delete(db.tasks, taskName(task))
	return nil

}
