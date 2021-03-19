package main

import (
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
}

func newStore() habitDBStore {
	return habitDBStore{
		habits: make(map[habitName]habit),
	}
}

func (db *habitDBStore) addHabit(newHabit habit) error {

	if _, exists := db.habits[habitName(newHabit.name)]; exists {
		return ErrHabitAlreadyExists
	}

	newHabit.adoptionRate = 0
	newHabit.createdAt = time.Now()

	db.habits[habitName(newHabit.name)] = newHabit
	return nil
}

func (db *habitDBStore) removeHabit(habit string) error {

	if _, exists := db.habits[habitName(habit)]; !exists {
		return ErrHabitNotExists
	}

	delete(db.habits, habitName(habit))

	return nil

}

func (db *habitDBStore) completeHabit(habit string) error {

	currentHabit, exists := db.habits[habitName(habit)]
	if !exists {
		return ErrHabitNotExists
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

func (db *habitDBStore) reportHabit(habit string) error {

	currentHabit, exists := db.habits[habitName(habit)]
	if !exists {
		return ErrHabitNotExists
	}

	fmt.Printf("Habit : %v\n", currentHabit.name)

	if currentHabit.description != "" {
		fmt.Printf("Description : %v\n", currentHabit.description)
	}

	fmt.Printf("\nAdoption Rate : %v\n", currentHabit.adoptionRate)
	fmt.Printf("Created At : %v\n", currentHabit.createdAt.String())
	fmt.Printf("Current Streak : %v\n", currentHabit.streak)

	if currentHabit.lastCompletionAt.IsZero() {
		fmt.Println("Not Started yet")
	} else {
		fmt.Printf("Last Completion at : %v\n", currentHabit.lastCompletionAt.String())
	}

	return nil
}
