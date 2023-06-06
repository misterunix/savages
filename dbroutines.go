package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/glebarez/go-sqlite"
	"github.com/misterunix/sniffle/hashing"
)

// Create a new DB. Remove the old one if it exists.
func CreateDB() {

	DropTable(SAVAGETABLE)
	DropTable(GAMEDBTABLE)
	DropTable(BIRTHRECORDTABLE)
	DropTable(LOGGINGTABLE)
	DropTable(USERSTABLE)

	tx := dbx.MustBegin()
	s := CreateTableFromStruct(SAVAGETABLE, Sav{})
	s = strings.ToLower(s)
	tx.MustExec(s)

	s = CreateTableFromStruct(GAMEDBTABLE, gamedb{})
	s = strings.ToLower(s)
	tx.MustExec(s)

	s = CreateTableFromStruct(BIRTHRECORDTABLE, birthrecord{})
	s = strings.ToLower(s)
	tx.MustExec(s)

	s = CreateTableFromStruct(LOGGINGTABLE, log{})
	s = strings.ToLower(s)
	tx.MustExec(s)

	s = CreateTableFromStruct(USERSTABLE, user{})
	s = strings.ToLower(s)
	tx.MustExec(s)

	tx.Commit()

	initGame()

	fmt.Println("Created a new database.")
	os.Exit(0)
}

// Drop a table
func DropTable(table string) {
	tx := dbx.MustBegin()
	s := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	tx.MustExec(s)
	tx.Commit()
}

// add the admin user with password
func addInitialUser() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Admin username: ")
	admin, err := reader.ReadString('\n')
	CheckErr(err, true)
	fmt.Print("Admin password: ")
	pwd, err := reader.ReadString('\n')
	CheckErr(err, true)

	u := user{}
	u.ID = 0
	u.Username = admin
	u.Email = "admin@localhost"
	u.Password = hashing.StringHash(hashing.SHA256, pwd)

	admin = strings.TrimRight(admin, "\n")
	pwd = strings.TrimRight(pwd, "\n")

	s := InsertIntoTable(USERSTABLE, u)

	tx := dbx.MustBegin()
	tx.MustExec(s)
	tx.Commit()
	fmt.Printf("user '%s' pwd '%s'\n", admin, pwd)
}

// setup the game day
func setGameDB() {
	g := gamedb{}
	g.ID = 0
	g.Day = 0
	s := InsertIntoTable(GAMEDBTABLE, g)
	tx := dbx.MustBegin()
	tx.MustExec(s)
	tx.Commit()
}

// Initialize the game.
func initGame() {
	setGameDB()

	addInitialUser()
}

// Add the starting savages.
func addStartingSavages() {
	tx := dbx.MustBegin()
	var s string

	for i := 0; i < gen0Count; i++ {
		g := Sav{}
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
		g.Strength = uint8(rnd.Intn(18))
		g.Intelligence = uint8(rnd.Intn(18))
		g.Charisma = uint8(rnd.Intn(18))
		g.Wisdom = uint8(rnd.Intn(18))
		g.Dexterity = uint8(rnd.Intn(18))
		g.Constitution = uint8(rnd.Intn(18))
		g.Hunger = g.HungerMax
		g.Thirst = g.ThirstMax
		g.Health = g.HealthMax
		g.Sex = uint8(rnd.Int() % 2)
		g.Pregnant = -1
		s = InsertIntoTable(SAVAGETABLE, g)
		tx.MustExec(s)
	}
	tx.Commit()
}
