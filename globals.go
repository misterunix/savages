package main

import (
	"database/sql"
	"math/rand"
	"os"
)

const VERSION = "0.0.0a"

const (
	maxX = 100
	maxY = 100
)

const gen0Count = 200

var database *sql.DB
var rnd *rand.Rand
var savs []Sav
var lastnames []string
var girlnames []string
var boynames []string
var logfile *os.File

const (
	SAVAGETABLE      = "sav"
	GAMEDBTABLE      = "gamedb"
	BIRTHRECORDTABLE = "birthrecords"
	LOGGINGTABLE     = "logging"
	USERSTABLE       = "users"
)
