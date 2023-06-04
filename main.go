package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	fmt.Println("Starting Savages.")
	fmt.Println("Version:", VERSION)

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	var err error

	createnewdb := false
	generation0 := false

	flag.BoolVar(&createnewdb, "newdb", false, "Create a new database.")
	flag.BoolVar(&generation0, "gen0", false, "Generate a new generation 0.")
	flag.Parse()

	if createnewdb {
		fmt.Println("Creating a new database.")
		database, err := sql.Open("sqlite", "db/savages.db")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer database.Close()

		DropTable(database, "savage")
		DropTable(database, "gamedb")
		DropTable(database, "birthrecords")
		DropTable(database, "logging")
		DropTable(database, "users")

		CreateDB(database)
	}

	if generation0 {
		fmt.Println("Generating a new generation 0.")
		database, err := sql.Open("sqlite", "db/savages.db")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer database.Close()

		o := "BEGIN;\n"
		beginstatement, err := database.Prepare(o)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = beginstatement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < gen0Count; i++ {
			g := savage{}
			g.ID = i
			g.OwnerID = 0
			g.Updated = false
			g.FirstName = "Gen"
			g.LastName = "Zero"
			g.Location = XY2Index(rnd.Intn(maxX), rnd.Intn(maxY))
			g.Age = 0
			g.FatherID = 0
			g.MotherID = 0
			g.HungerMax = uint8(rnd.Intn(50) + 50)
			g.ThirstMax = uint8(rnd.Intn(50) + 50)
			g.HealthMax = uint8(rnd.Intn(50) + 50)
			g.Strength = uint8(rnd.Intn(22) + 3)
			g.Intelligence = uint8(rnd.Intn(22) + 3)
			g.Charisma = uint8(rnd.Intn(22) + 3)
			g.Wisdom = uint8(rnd.Intn(22) + 3)
			g.Dexterity = uint8(rnd.Intn(22) + 3)
			g.Constitution = uint8(rnd.Intn(22) + 3)
			g.Hunger = g.HungerMax
			g.Thirst = g.ThirstMax
			g.Health = g.HealthMax
			g.Sex = uint8(rnd.Int() % 2)
			g.Pregnant = -1
			s := InsertIntoTable("savage", g)
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
		o = "COMMIT;\n"
		beginstatement, err = database.Prepare(o)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = beginstatement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Generated a new generation 0.")
		os.Exit(0)
	}

	database, err = sql.Open("sqlite", "db/savages.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer database.Close()

	RunDay()

}

func RunDay() {
	fmt.Println("Running a day.")
	/*
		database, err := sql.Open("sqlite", "db/savages.db")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer database.Close()
	*/
	countSQL := "SELECT COUNT(*) FROM savage WHERE health > 0;"
	statement, err := database.Prepare(countSQL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var c int // count of alive savages
	for rows.Next() {
		rows.Scan(&c)
	}
	rows.Close()

	// Load the savages into memory.
	//savages := make([]savage, c)
	savagesSQL := "SELECT * FROM savage WHERE health > 0;"
	statement, err = database.Prepare(savagesSQL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, err = statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; rows.Next(); i++ {
		ss := savage{}

		rows.Scan(
			&ss.ID,
			&ss.OwnerID,
			&ss.Updated,
			&ss.Location,
			&ss.FirstName,
			&ss.LastName,
			&ss.Age,
			&ss.Sex,
			&ss.Pregnant,
			&ss.MotherID,
			&ss.FatherID,
			&ss.Hunger,
			&ss.HungerMax,
			&ss.Thirst,
			&ss.ThirstMax,
			&ss.Health,
			&ss.HealthMax,
			&ss.Strength,
			&ss.Intelligence,
			&ss.Charisma,
			&ss.Wisdom,
			&ss.Dexterity,
			&ss.Constitution,
		)

		savages = append(savages, ss)
	}
	rows.Close()

	for i := 0; i < c; i++ {
		savages[i].Updated = true
		savages[i].Age++
	}

	// Load the distances into memory.
	var distances []distance

	//distances := make([]distance, c*c)
	for i := 0; i < c-1; i++ {
		for j := i + 1; j < c; j++ {
			d := distance{}
			d.ID1 = savages[i].ID
			d.ID2 = savages[j].ID
			d.Distance = DistanceSavage(savages[i], savages[j])
			distances = append(distances, d)
		}
	}

	/*
		for _, s := range savages {
			fmt.Println(s)
		}
	*/

	//
	// Age the savages.
	//
	o := "BEGIN;\n"
	beginstatement, err := database.Prepare(o)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = beginstatement.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < c; i++ {
		id := savages[i].ID
		sql1 := "UPDATE savage SET age = age + 1 WHERE id = '" + fmt.Sprintf("%d", id) + "';"
		statement, err := database.Prepare(sql1)
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
	o = "COMMIT;\n"
	beginstatement, err = database.Prepare(o)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = beginstatement.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//
	//
	//

	fmt.Println("Distances")
	for _, d := range distances {
		if d.Distance < 11 {
			fmt.Println(d)
			if savages[d.ID1].Sex != savages[d.ID2].Sex {
				fmt.Println("Breed?")
				if savages[d.ID1].Pregnant != -1 || savages[d.ID2].Pregnant != -1 {
					// someone is already pregnant
					continue
				}

				if savages[d.ID1].Age >= 14 && savages[d.ID2].Age >= 14 {
					breed(d.ID1, d.ID2)
				}

				//fmt.Println("Next savage id:", getNextSavageID())
			}
		}
	}

}

// Get the next savage id from the database
func getNextSavageID() int {
	var lastID int

	sql1 := "SELECT MAX(ID) FROM savage;"
	statement, err := database.Prepare(sql1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, err := statement.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; rows.Next(); i++ {
		rows.Scan(&lastID)
	}
	rows.Close()
	lastID++
	return lastID
}
