package main

import (
	"fmt"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/misterunix/sniffle/hashing"
)

// Create a new DB
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

	//fmt.Println(schema)
	//dbx.MustExec(schema)

	/*

		fmt.Println(s)
		statement, err := d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		s = CreateTableFromStruct("gamedb", gamedb{})
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		s = CreateTableFromStruct("birthrecords", birthrecord{})
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		s = CreateTableFromStruct("logging", log{})
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		s = CreateTableFromStruct("users", user{})
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		u := user{}
		u.ID = 0
		u.Username = "admin"
		u.Email = "admin@localhost"
		u.Password = hashing.StringHash(hashing.SHA256, "DefaultFuckingPassword")
		s = InsertIntoTable("users", u)
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()

		g := gamedb{}
		g.ID = 0
		g.Day = 0
		s = InsertIntoTable("gamedb", g)
		fmt.Println(s)
		statement, err = d.Prepare(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = statement.Exec()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statement.Close()
	*/
	p := hashing.StringHash(hashing.SHA256, "DefaultFuckingPassword")
	fmt.Println(p)
	fmt.Println("Created a new database.")
	os.Exit(0)
}

// Remove any of the tables in the database.
func DropTable(table string) {
	tx := dbx.MustBegin()
	s := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	tx.MustExec(s)
	tx.Commit()
	/*
	   //fmt.Println("Drop table:", table)
	   s := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	   fmt.Println(s)
	   statement, _ := d.Prepare(s)
	   _, err := statement.Exec()
	   CheckErr(err, true)
	*/
}
