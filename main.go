package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("Starting Savages.")
	fmt.Println("Version:", VERSION)

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	//var err error

	createnewdb := false
	generation0 := false

	flag.BoolVar(&createnewdb, "newdb", false, "Create a new database.")
	flag.BoolVar(&generation0, "gen0", false, "Generate a new generation 0.")
	flag.Parse()

	var err error
	dbx, err = sqlx.Connect("sqlite", "db/savages.db")
	_ = CheckErr(err, true)
	defer dbx.Close()

	if createnewdb {
		fmt.Println("Creating a new database.")
		CreateDB()
		fmt.Println("Created a new database.")
	}

	if generation0 {
		fmt.Println("Generating a new generation 0.")
		addStartingSavages()
		fmt.Println("Generated a new generation 0.")
		os.Exit(0)
	}

	//var err error
	dbx, err = sqlx.Connect("sqlite", "db/savages.db")
	_ = CheckErr(err, true)
	defer dbx.Close()

	tmpString := fmt.Sprintf("SELECT day FROM %s WHERE ID='0';", GAMEDBTABLE)
	//dayNumSQL := "SELECT day FROM gamedb WHERE ID='0';"
	var dayNum int
	err = dbx.Get(&dayNum, tmpString)
	CheckErr(err, true)
	dayNum++
	fmt.Println("Day:", dayNum)

	tmpString = fmt.Sprintf("UPDATE %s SET day = day + 1 WHERE ID='0';", GAMEDBTABLE)
	//s := "UPDATE gamedb SET day = day + 1 WHERE ID='0';"
	tx := dbx.MustBegin()
	tx.MustExec(tmpString)
	tx.Commit()

	/*
		database, err = sql.Open("sqlite", "db/savages.db")
		_ = CheckErr(err, true)
		defer database.Close()
	*/

	RunDay()

}

func RunDay() {
	var count int // count of alive savages
	fmt.Println("Running a day.")
	/*
		database, err := sql.Open("sqlite", "db/savages.db")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer database.Close()
	*/

	// Increment the day.
	tmpString := fmt.Sprintf("UPDATE %s SET day = day + 1 WHERE ID='0';", GAMEDBTABLE)
	//s := "UPDATE gamedb SET day = day + 1 WHERE ID='0';"
	tx := dbx.MustBegin()
	tx.MustExec(tmpString)
	tx.Commit()

	tmpString = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE health > 0;", SAVAGETABLE)
	err := dbx.Get(&count, tmpString)
	_ = CheckErr(err, true)

	fmt.Println("There are", count, "savages alive.")

	s := fmt.Sprintf("SELECT * FROM %s WHERE health > 0;", SAVAGETABLE)
	rows1, err := dbx.Queryx(s)
	_ = CheckErr(err, true)
	for rows1.Next() {
		var ss Sav
		err := rows1.StructScan(&ss)
		_ = CheckErr(err, true)
		fmt.Println(ss)
		savs = append(savs, ss)
	}

	//dbx.Select(&savs, s)
	fmt.Println("dbloaded", len(savs))
	/*
		rows, err := dbx.Queryx(s)
		_ = CheckErr(err, true)
		for rows.Next() {
			err := rows.StructScan(&ss)
			_ = CheckErr(err, true)
			savs = append(savs, ss)
		}
	*/
	fmt.Println("Completed")
	for _, j := range savs {
		fmt.Println(j)
	}
	//err = dbx.Select(&savages, "SELECT * FROM savage WHERE health > 0;")
	//CheckErr(err, true)

	// Load the distances into memory.
	var distances []distance

	//distances = make([]distance, count*count)

	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			d := distance{}
			d.ID1 = savs[i].ID
			d.ID2 = savs[j].ID
			d.Distance = DistanceSavage(savs[i], savs[j])
			if d.Distance <= 10 {
				distances = append(distances, d)
			}
		}
	}

	fmt.Println("There are", len(distances), "distances.")

	/*
		for _, s := range savages {
			fmt.Println(s)
		}
	*/

	//
	// Age the savages.
	//
	/*
		o := "BEGIN;\n"
		beginstatement, err := database.Prepare(o)
		_ = CheckErr(err, true)
		_, err = beginstatement.Exec()
		_ = CheckErr(err, true)
		for i := 0; i < c; i++ {
			id := savages[i].ID
			sql1 := "UPDATE savage SET age = age + 1 WHERE id = '" + fmt.Sprintf("%d", id) + "';"
			statement, err := database.Prepare(sql1)
			_ = CheckErr(err, true)
			_, err = statement.Exec()
			_ = CheckErr(err, true)
		}
		o = "COMMIT;\n"
		beginstatement, err = database.Prepare(o)
		_ = CheckErr(err, true)
		_, err = beginstatement.Exec()
		_ = CheckErr(err, true)

		//
		//
		//

		fmt.Println("Distances")
		for _, d := range distances {
			if d.Distance < 11 {
				//fmt.Println(d)
				if savages[d.ID1].Sex != savages[d.ID2].Sex {
					//fmt.Println("Breed?")
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
	*/
}

// Get the next savage id from the database
func getNextSavageID() int {
	var lastID int

	sql1 := fmt.Sprintf("SELECT MAX(ID) FROM savage;", SAVAGETABLE)
	statement, err := database.Prepare(sql1)
	_ = CheckErr(err, true)
	rows, err := statement.Query()
	_ = CheckErr(err, true)
	for i := 0; rows.Next(); i++ {
		rows.Scan(&lastID)
	}
	rows.Close()
	lastID++
	return lastID
}
