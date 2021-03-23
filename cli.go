package main

import (
	"flag"
	"fmt"
)

func main() {
	createJSONFile(jsonFileName)
	db := newJSONStore()
	var err error

	option := flag.String("option", "", "habit")
	habitName := flag.String("name", "", "name for habit")
	habitDesc := flag.String("desc", "", "description for habit")

	flag.Parse()

	if *option == "" {
		fmt.Println("Habit-Tracker")
		return
	}

	if *option == "new" {
		if *habitName == "" {
			fmt.Println("Add name of habit")
			return
		}
		if *habitDesc == "" {
			fmt.Println("Add description of habit")
			return
		}

		err = db.addHabit(habit{Name: *habitName, Description: *habitDesc})
		checkErr(err)
	}

	if *option == "remove" {
		if *habitName == "" {
			fmt.Println("Add name of habit")
			return
		}
		err = db.removeHabit(*habitName)
		checkErr(err)
	}

	if *option == "complete" {
		if *habitName == "" {
			fmt.Println("Add name of habit")
			return
		}
		err = db.completeHabit(*habitName)
		checkErr(err)

		fmt.Printf("Completed %s\n", *habitName)
	}

	if *option == "report" {
		if *habitName == "" {
			fmt.Println("Add name of habit")
			return
		}
		err = db.reportHabit(*habitName)
		checkErr(err)
	}
}
