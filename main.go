package main

import (
	"flag"
	"fmt"
	"log"

	"os"

	"github.com/blockloop/scan"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	fmt.Println("Starting Savages.")
	fmt.Println("Version:", VERSION)

	setupLog()
	//log.SetOutput(os.Stdout)
	//log.SetFlags(log.Llongfile)

	//rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	//r := rand.IntN(100)
	//var err error

	createnewdb := false
	generation0 := false

	flag.BoolVar(&createnewdb, "newdb", false, "Create a new database.")
	flag.BoolVar(&generation0, "gen0", false, "Generate a new generation 0.")
	flag.Parse()

	err := OpenDB() // Open the database
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//defer database.Close()

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

	loadBoyNames()
	loadGirlNames()
	loadLastNames()

	for i := 0; i < 10; i++ {

		tmpString := fmt.Sprintf("SELECT day FROM %s WHERE ID='0';", GAMEDBTABLE)
		//dayNumSQL := "SELECT day FROM gamedb WHERE ID='0';"
		var dayNum int
		database.QueryRow(tmpString).Scan(&dayNum)
		fmt.Println("Day:", dayNum)

		dayNum++
		fmt.Println("Day:", dayNum)

		/*
			tmpString = fmt.Sprintf("UPDATE %s SET day = day + 1 WHERE ID='0';", GAMEDBTABLE)
			//s := "UPDATE gamedb SET day = day + 1 WHERE ID='0';"
			tx := dbx.MustBegin()
			tx.MustExec(tmpString)
			tx.Commit()
		*/
		/*
			database, err = sql.Open("sqlite", "db/savages.db")
			_ = CheckErr(err, true)
			defer database.Close()
		*/

		savs = make([]Sav, 0)
		err := RunDay()
		if err != nil {
			log.Println(err)
			cleanExit(1)
		}
	}
	cleanExit(0)
}

func setupLog() error {
	var err error
	logfile, err = os.OpenFile("savages.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	return nil
}

func cleanExit(exitcode int) {
	log.Println("Exiting.")

	err := database.Close()
	if err != nil {
		log.Println("Error closing database:", err)
	}

	err = logfile.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error closing logfile:", err)
		os.Exit(1)
	}

	os.Exit(exitcode)
}

func RunDay() error {
	var count int // count of alive savages
	fmt.Println("Running a day.")

	// Increment the day.
	tmpString := fmt.Sprintf("UPDATE %s SET day = day + 1 WHERE ID='0';", GAMEDBTABLE)
	//s := "UPDATE gamedb SET day = day + 1 WHERE ID='0';"
	statement, err := database.Prepare(tmpString)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	tmpString = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE health > 0;", SAVAGETABLE)
	err = database.QueryRow(tmpString).Scan(&count)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("There are", count, "savages alive.")

	s := fmt.Sprintf("SELECT * FROM %s WHERE health > 0;", SAVAGETABLE)
	rows, err := database.Query(s)
	if err != nil {
		log.Println(err)
		return err
	}

	err = scan.Rows(&savs, rows)
	if err != nil {
		log.Println(err)
		return err
	}
	// for rows.Next() {
	// 	//var ss Sav
	// 	//err := rows.Scan(&ss)
	// 	err := scan.Rows(&savs, rows)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	//fmt.Println(cols)
	// 	//savs = append(savs, ss)
	// }

	// for i, j := range savs {
	// 	fmt.Println(i, j)
	// }

	fmt.Println("dbloaded", len(savs))

	fmt.Println("Completed")
	// for _, j := range savs {
	// 	fmt.Println(j)
	// }

	var distances []distance

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

	return nil
}

// Get the next savage id from the database
func getNextSavageID() int {
	var lastID int

	sql1 := fmt.Sprintf("SELECT MAX(ID) FROM %s;", SAVAGETABLE)
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
