package main

import (
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func ageSavages() error {

	sqlstring := "update savs set age = age + 1 where Health> 0;"
	_, err := database.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getCountOfSavages() (int, error) {
	var count int
	err := database.QueryRow("select count(*) from savs where Health > 0;").Scan(&count)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
