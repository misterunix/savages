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

	DropTable("savage")
	DropTable("gamedb")
	DropTable("birthrecords")
	DropTable("logging")
	DropTable("users")

	tx := dbx.MustBegin()
	s := CreateTableFromStruct("savage", savage{})
	tx.MustExec(s)
	s = CreateTableFromStruct("gamedb", gamedb{})
	tx.MustExec(s)
	s = CreateTableFromStruct("birthrecords", birthrecord{})
	tx.MustExec(s)
	s = CreateTableFromStruct("logging", log{})
	tx.MustExec(s)
	s = CreateTableFromStruct("users", user{})
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

	s := InsertIntoTable("users", u)

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
	s := InsertIntoTable("gamedb", g)
	tx := dbx.MustBegin()
	tx.MustExec(s)
	tx.Commit()
}

// Initialize the game.
func initGame() {
	setGameDB()

	addInitialUser()
}
