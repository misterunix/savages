package main

func mating(id1, id2 int) {

	i1 := (savs[id1].Charisma +
		savs[id1].Intelligence +
		savs[id1].Strength +
		savs[id1].Wisdom +
		savs[id1].Dexterity +
		savs[id1].Constitution) / 6

	i2 := (savs[id2].Charisma +
		savs[id2].Intelligence +
		savs[id2].Strength +
		savs[id2].Wisdom +
		savs[id2].Dexterity +
		savs[id2].Constitution) / 6

	var v float64
	if i1 > i2 {
		v = float64(i2) / float64(i1)
	} else {
		v = float64(i1) / float64(i2)
	}

	if v > 0.75 {
		breed(id1, id2)
	}

}

func SetPregnant(id1, id2 int) {
	if savs[id1].Sex == 1 {
		savs[id1].Pregnant = 10
		savs[id1].Updated = true
		savs[id1].Points += 2
	} else {
		savs[id2].Pregnant = 10
		savs[id2].Updated = true
		savs[id2].Points++
	}
}

func breed(id1, id2 int) {

	newSavage := Sav{}

	newSavage.ID = getNextSavageID()

	newSavage.Updated = false

	if rnd.Intn(100) <= 50 {
		newSavage.Location = savs[id1].Location
	} else {
		newSavage.Location = savs[id2].Location
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
		newSavage.OwnerID = savs[id1].OwnerID
	} else {
		newSavage.OwnerID = savs[id2].OwnerID
	}

	if rnd.Intn(100) <= 50 {
		newSavage.MotherID = savs[id1].ID
		newSavage.FatherID = savs[id2].ID
	} else {
		newSavage.MotherID = savs[id2].ID
		newSavage.FatherID = savs[id1].ID
	}

	if rnd.Intn(100) <= 50 {
		newSavage.HungerMax = savs[id1].HungerMax
	} else {
		newSavage.HungerMax = savs[id2].HungerMax
	}

	if rnd.Intn(100) <= 50 {
		newSavage.ThirstMax = savs[id1].ThirstMax
	} else {
		newSavage.ThirstMax = savs[id2].ThirstMax
	}

	if rnd.Intn(100) <= 50 {
		newSavage.HealthMax = savs[id1].HealthMax
	} else {
		newSavage.HealthMax = savs[id2].HealthMax
	}

	newSavage.Hunger = newSavage.HungerMax
	newSavage.Thirst = newSavage.ThirstMax
	newSavage.Health = newSavage.HealthMax

	if rnd.Intn(100) <= 50 {
		newSavage.Strength = savs[id1].Strength
	} else {
		newSavage.Strength = savs[id2].Strength
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Intelligence = savs[id1].Intelligence
	} else {
		newSavage.Intelligence = savs[id2].Intelligence
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Charisma = savs[id1].Charisma
	} else {
		newSavage.Charisma = savs[id2].Charisma
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Wisdom = savs[id1].Wisdom
	} else {
		newSavage.Wisdom = savs[id2].Wisdom
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Dexterity = savs[id1].Dexterity
	} else {
		newSavage.Dexterity = savs[id2].Dexterity
	}

	if rnd.Intn(100) <= 50 {
		newSavage.Constitution = savs[id1].Constitution
	} else {
		newSavage.Constitution = savs[id2].Constitution
	}

	s := InsertIntoTable(SAVAGETABLE, newSavage)
	statement, err := database.Prepare(s)
	_ = CheckErr(err, true)
	_, err = statement.Exec()
	_ = CheckErr(err, true)

}
