package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_newStore(t *testing.T) {
	tests := []struct {
		name string
		want habitDBStore
	}{
		{"check new store", habitDBStore{habits: make(map[habitName]habit)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStore() do not create habitDBStore")
			}
		})
	}
}

func Test_habitDBStore_addHabit(t *testing.T) {

	db := newStore()

	tests := []struct {
		name    string
		habit   habit
		wantErr error
	}{
		{"Adding 1st habit", habit{name: "Exercise"}, nil},
		{"Adding habit which already exists", habit{name: "Exercise"}, ErrHabitAlreadyExists},
		{"Adding 2nd habit", habit{name: "Reading"}, nil},
	}
	for _, tt := range tests {
		if err := db.addHabit(tt.habit); err != tt.wantErr {
			t.Errorf("AddHabit failed at %v:  got = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}

}

func Test_habitDBStore_removeHabit(t *testing.T) {

	db := newStore()
	db.addHabit(habit{name: "Exercise"})

	tests := []struct {
		name    string
		habit   string
		wantErr error
	}{
		{"Removing already existing habit", "Exercise", nil},
		{"Removing habit that's not added", "Exercise", ErrHabitNotExists},
	}
	for _, tt := range tests {
		if err := db.removeHabit(tt.habit); err != tt.wantErr {
			t.Errorf("Remove Habit failed at %v:  got = %v, wantErr %v", tt.name, err, tt.wantErr)
		}

	}
}

func Test_habitDBStore_completeHabit(t *testing.T) {
	db := newStore()

	db.addHabit(habit{name: "Exercise"})

	tests := []struct {
		name        string
		habit       string
		streak      int
		completedAt time.Time
		wantErr     error
	}{
		{"Completing existing habit", "Exercise", 1, time.Now(), nil},
		{"Completing a non existing habit", "Read", 0, time.Now(), ErrHabitNotExists},
	}
	for _, tt := range tests {
		if err := db.completeHabit(tt.habit); err != tt.wantErr {
			t.Errorf("Completing a habit failed: got = %v, wantErr %v", err, tt.wantErr)
		}

		hbt, present := db.habits[habitName(tt.habit)]

		if present && hbt.streak != tt.streak {
			t.Errorf("Streak attribute failed: got = %v, wantErr %v", hbt.streak, tt.streak)
		}

	}
}
