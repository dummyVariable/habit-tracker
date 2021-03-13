package main

import (
	"errors"
	"fmt"
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

func (db *habitDBStore) completeHabit(habit string) error {

	currentHabit, exists := db.habits[habitName(habit)]
	if !exists {
		return errors.New("Habit not exists")
	}
	currentTime := time.Now()

	if currentHabit.lastCompletionAt.Sub(currentTime).Hours() < 25 {
		currentHabit.streak++
	} else {
		currentHabit.streak = 0
	}

	currentHabit.lastCompletionAt = currentTime
	db.habits[habitName(habit)] = currentHabit

	return nil
}

func (db *habitDBStore) completeTask(habit, task string, reps int) error {

	_, exists := db.habits[habitName(habit)]
	if !exists {
		return errors.New("Habit not exists")
	}

	currentTask, exists := db.tasks[taskName(task)]
	if !exists {
		return errors.New("Task not exists")
	}

	taskList := db.habitTaskMap[habitName(habit)]

	if !contains(taskList, taskName(task)) {
		return errors.New("Task does not relate to habit")
	}

	currentTime := time.Now()

	currentTask.reps = reps
	currentTask.completedAt = currentTime

	if currentTask.bestReps < reps {
		currentTask.bestReps = reps
	}

	db.tasks[taskName(task)] = currentTask

	return nil
}
func (db *habitDBStore) reportHabit(habit string) error {

	currentHabit, exists := db.habits[habitName(habit)]
	if !exists {
		return errors.New("Habit not exists")
	}

	taskList := db.habitTaskMap[habitName(habit)]

	fmt.Println(currentHabit.name)

	for i, task := range taskList {
		currentTask := db.tasks[task]
		fmt.Println(i+1, currentTask.name, currentTask.bestReps)
	}

	fmt.Println(currentHabit.streak)

	if currentHabit.lastCompletionAt.IsZero() {
		fmt.Println("Not started yet")
	} else {
		fmt.Println(currentHabit.lastCompletionAt.Date())
	}

	return nil
}
