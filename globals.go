package main

import (
	"database/sql"
	"math/rand"

	"github.com/jmoiron/sqlx"
)

const VERSION = "0.0.0a"

const (
	maxX = 100
	maxY = 100
)

const gen0Count = 100

var database *sql.DB
var dbx *sqlx.DB

var rnd *rand.Rand
var savs []Sav

const (
	SAVAGETABLE      = "sav"
	GAMEDBTABLE      = "gamedb"
	BIRTHRECORDTABLE = "birthrecords"
	LOGGINGTABLE     = "logging"
	USERSTABLE       = "users"
)
