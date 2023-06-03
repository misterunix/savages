package main

import (
	"fmt"
	"os"
)

func breed(id1, id2 int) {

	newSavage := savage{}

	newSavage.ID = getNextSavageID()

	newSavage.Update = false

	if rnd.Intn(100) <= 50 {
		newSavage.Location = savages[id1].Location
	} else {
		newSavage.Location = savages[id2].Location
	}

	newSavage.FirstName = "Gen"
	newSavage.LastName = "B"
	newSavage.Age = 0

	if rnd.Intn(100) <= 50 {
		newSavage.Sex = 0
	} else {
		newSavage.Sex = 1
	}

	if rnd.Intn(100) <= 50 {
		newSavage.MotherID = savages[id1].ID
		newSavage.FatherID = savages[id2].ID
	} else {
		newSavage.MotherID = savages[id2].ID
		newSavage.FatherID = savages[id1].ID
	}

	if rnd.Intn(100) <= 50 {
		newSavage.HungerMax = savages[id1].HungerMax
	} else {
		newSavage.HungerMax = savages[id2].HungerMax
	}

	if rnd.Intn(100) <= 50 {
		newSavage.ThirstMax = savages[id1].ThirstMax
	} else {
		newSavage.ThirstMax = savages[id2].ThirstMax
	}

	if rnd.Intn(100) <= 50 {
		newSavage.HealthMax = savages[id1].HealthMax
	} else {
		newSavage.HealthMax = savages[id2].HealthMax
	}

	newSavage.Hunger = newSavage.HungerMax
	newSavage.Thirst = newSavage.ThirstMax
	newSavage.Health = newSavage.HealthMax

	if rnd.Intn(100) <= 50 {
		newSavage.Strength = savages[id1].Strength
	} else {
		newSavage.Strength = savages[id2].Strength
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Intelligence = savages[id1].Intelligence
	} else {
		newSavage.Intelligence = savages[id2].Intelligence
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Charisma = savages[id1].Charisma
	} else {
		newSavage.Charisma = savages[id2].Charisma
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Wisdom = savages[id1].Wisdom
	} else {
		newSavage.Wisdom = savages[id2].Wisdom
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Dexterity = savages[id1].Dexterity
	} else {
		newSavage.Dexterity = savages[id2].Dexterity
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Constitution = savages[id1].Constitution
	} else {
		newSavage.Constitution = savages[id2].Constitution
	}

	s := InsertIntoTable("savage", newSavage)
	statement, err := database.Prepare(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
