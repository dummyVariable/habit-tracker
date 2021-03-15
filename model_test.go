package main

import (
	"reflect"
	"testing"
)

func Test_newStore(t *testing.T) {
	tests := []struct {
		name string
		want habitDBStore
	}{
		{"check new store", habitDBStore{habits: make(map[habitName]habit), tasks: make(map[taskName]task), habitTaskMap: make(map[habitName][]taskName)}},
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

func Test_habitDBStore_addTask(t *testing.T) {
	type fields struct {
		habits       map[habitName]habit
		tasks        map[taskName]task
		habitTaskMap map[habitName][]taskName
	}
	type args struct {
		habit   string
		newTask task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &habitDBStore{
				habits:       tt.fields.habits,
				tasks:        tt.fields.tasks,
				habitTaskMap: tt.fields.habitTaskMap,
			}
			if err := db.addTask(tt.args.habit, tt.args.newTask); (err != nil) != tt.wantErr {
				t.Errorf("habitDBStore.addTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_habitDBStore_removeTask(t *testing.T) {
	type fields struct {
		habits       map[habitName]habit
		tasks        map[taskName]task
		habitTaskMap map[habitName][]taskName
	}
	type args struct {
		habit string
		task  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &habitDBStore{
				habits:       tt.fields.habits,
				tasks:        tt.fields.tasks,
				habitTaskMap: tt.fields.habitTaskMap,
			}
			if err := db.removeTask(tt.args.habit, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("habitDBStore.removeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_habitDBStore_completeHabit(t *testing.T) {
	type fields struct {
		habits       map[habitName]habit
		tasks        map[taskName]task
		habitTaskMap map[habitName][]taskName
	}
	type args struct {
		habit string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &habitDBStore{
				habits:       tt.fields.habits,
				tasks:        tt.fields.tasks,
				habitTaskMap: tt.fields.habitTaskMap,
			}
			if err := db.completeHabit(tt.args.habit); (err != nil) != tt.wantErr {
				t.Errorf("habitDBStore.completeHabit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_habitDBStore_completeTask(t *testing.T) {
	type fields struct {
		habits       map[habitName]habit
		tasks        map[taskName]task
		habitTaskMap map[habitName][]taskName
	}
	type args struct {
		habit string
		task  string
		reps  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &habitDBStore{
				habits:       tt.fields.habits,
				tasks:        tt.fields.tasks,
				habitTaskMap: tt.fields.habitTaskMap,
			}
			if err := db.completeTask(tt.args.habit, tt.args.task, tt.args.reps); (err != nil) != tt.wantErr {
				t.Errorf("habitDBStore.completeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
