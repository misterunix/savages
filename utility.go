package main

import (
	"fmt"
	"os"
)

func CheckErr(err error, die bool) error {
	if err != nil {
		fmt.Println(err)
		if die {
			os.Exit(1)
		} else {
			return err
		}
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
