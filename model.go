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

	if _, exists := db.habits[habitName(newHabit.Name)]; exists {
		return ErrHabitAlreadyExists
	}

	newHabit.AdoptionRate = 0
	newHabit.CreatedAt = time.Now()

	db.habits[habitName(newHabit.Name)] = newHabit
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

	if currentHabit.LastCompletionAt.Sub(currentTime).Hours() < 25 {
		currentHabit.Streak++
	} else {
		currentHabit.Streak = 0
	}

	currentHabit.LastCompletionAt = currentTime
	db.habits[habitName(habit)] = currentHabit

	return nil
}

func (db *habitDBStore) reportHabit(habit string) error {

	currentHabit, exists := db.habits[habitName(habit)]
	if !exists {
		return ErrHabitNotExists
	}

	fmt.Printf("Habit : %v\n", currentHabit.Name)

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
