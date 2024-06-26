package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	_ "github.com/glebarez/go-sqlite"

	"github.com/misterunix/sniffle/hashing"
)

// Open the database. If it doesn't exist, create it. Return an error if there is a problem.
func OpenDB() error {
	var err error
	fn := "db/savages.db"
	database, err = sql.Open("sqlite", fn)
	if err != nil {
		return err
	}
	database.SetMaxOpenConns(1)
	return nil
}

// Create a new DB. Remove the old one if it exists.
func CreateDB() error {

	DropTable(SAVAGETABLE)
	DropTable(GAMEDBTABLE)
	DropTable(BIRTHRECORDTABLE)
	DropTable(LOGGINGTABLE)
	DropTable(USERSTABLE)
	DropTable(SAVTOSAV)

	var s string
	s = "BEGIN;\n"
	s += CreateTableFromStruct(SAVAGETABLE, Sav{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	s += CreateTableFromStruct(GAMEDBTABLE, gamedb{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	s += CreateTableFromStruct(BIRTHRECORDTABLE, birthrecord{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	s += CreateTableFromStruct(LOGGINGTABLE, tlog{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	s += CreateTableFromStruct(USERSTABLE, user{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	s += CreateTableFromStruct(SAVTOSAV, savtosav{})
	s += "\n"

	s += "COMMIT;\n"
	fmt.Println(s)
	statement, err := database.Prepare(s)
	if err != nil {
		log.Println(err, s)
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	initGame()

	fmt.Println("Created a new database.")
	os.Exit(0)
	return nil
}

// Drop a table if it exists.
func DropTable(table string) {
	statement := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	database.Exec(statement)
}

// add the admin user with password
func addInitialUser() error {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Admin username: ")
	admin, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Print("Admin password: ")
	pwd, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return err
	}

	u := user{}
	u.ID = 0
	u.Username = admin
	u.Email = "admin@localhost"
	u.Password = hashing.StringHash(hashing.SHA256, pwd)

	admin = strings.TrimRight(admin, "\n")
	pwd = strings.TrimRight(pwd, "\n")

	s1 := InsertIntoTable(USERSTABLE, u)
	s := "BEGIN;\n" + s1 + "COMMIT;\n"
	statement, err := database.Prepare(s)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("user '%s' pwd '%s'\n", admin, pwd)
	return nil
}

// setup the game day
func setGameDB() error {
	g := gamedb{}
	g.ID = 0
	g.Day = 0
	s1 := InsertIntoTable(GAMEDBTABLE, g)
	s0 := "BEGIN;\n" + s1 + "COMMIT;\n"
	statemenr, err := database.Prepare(s0)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statemenr.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Initialize the game.
func initGame() {
	err := setGameDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = addInitialUser()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// get a last name. since its called a lot, open the file once and read it into a slice.
func loadLastNames() {
	lnf, err := os.OpenFile("lastnames.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	// var lastnames []string
	scanner := bufio.NewScanner(lnf)
	for scanner.Scan() {
		tmps := scanner.Text()
		tmps = strings.ToUpper(tmps[:1]) + tmps[1:]
		lastnames = append(lastnames, tmps)
	}
	err = lnf.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Last:", len(lastnames))
}

func loadGirlNames() {
	fnff, err := os.OpenFile("girls.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}

	scanner := bufio.NewScanner(fnff)
	for scanner.Scan() {
		girlnames = append(girlnames, scanner.Text())
	}
	err = fnff.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Girl:", len(girlnames))
}

func loadBoyNames() {

	fnmf, err := os.OpenFile("boys.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	scanner := bufio.NewScanner(fnmf)
	for scanner.Scan() {
		boynames = append(boynames, scanner.Text())
	}
	err = fnmf.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Boy:", len(boynames))
}

// Add the starting savages.
func addStartingSavages() error {

	// lnf, err := os.OpenFile("lastnames.txt", os.O_RDONLY, 0666)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// var lastnames []string
	// scanner := bufio.NewScanner(lnf)
	// for scanner.Scan() {
	// 	lastnames = append(lastnames, scanner.Text())
	// }
	// err = lnf.Close()
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	//lnc := len(lastnames)
	gnc := len(girlnames)
	bnc := len(boynames)
	lnc := len(lastnames)

	//fmt.Println(gnc, bnc, lnc)

	var s string
	var sqlog string
	s = "BEGIN;\n"
	sqlog = "BEGIN;\n"

	for i := 0; i < gen0Count; i++ {
		g := Sav{}
		g.ID = i
		g.OwnerID = 0
		g.Points = 0
		g.Updated = false
		//g.FirstName = "Gen"
		//g.LastName = "Zero"
		g.Generation = 0
		g.Location = XY2Index(rand.Intn(maxX), rand.Intn(maxY))
		g.Age = 0
		g.FatherID = 0
		g.MotherID = 0
		g.HungerMax = uint8(rand.Intn(50) + 50)
		g.ThirstMax = uint8(rand.Intn(50) + 50)
		g.HealthMax = uint8(rand.Intn(50) + 50)
		g.Strength = uint8(rand.Intn(17)) + 1
		g.Intelligence = uint8(rand.Intn(17)) + 1
		g.Charisma = uint8(rand.Intn(17)) + 1
		g.Wisdom = uint8(rand.Intn(17)) + 1
		g.Dexterity = uint8(rand.Intn(17)) + 1
		g.Constitution = uint8(rand.Intn(17)) + 1
		g.Hunger = g.HungerMax
		g.Thirst = g.ThirstMax
		g.Health = g.HealthMax
		g.Sex = uint8(rand.Int() % 2)

		g.LastName = lastnames[rand.Intn(lnc)]

		if g.Sex == 0 {
			g.FirstName = boynames[rand.Intn(bnc)]
		} else {
			g.FirstName = girlnames[rand.Intn(gnc)]
		}
		g.Pregnant = -1
		s += InsertIntoTable(SAVAGETABLE, g)

		tl := tlog{}
		tl.Date = 0
		tl.Who = 0
		tl.Message = fmt.Sprintf("Added %s %s to the game.", g.FirstName, g.LastName)
		sqlog += InsertIntoTable(LOGGINGTABLE, tl)

	}
	s += "COMMIT;\n"
	sqlog += "COMMIT;\n"

	statement, err := database.Prepare(s)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	statement, err = database.Prepare(sqlog)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
