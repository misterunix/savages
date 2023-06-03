package main

import (
	"database/sql"
	"math/rand"
)

const VERSION = "0.0.0a"

const (
	maxX = 100
	maxY = 100
)

const gen0Count = 100

var database *sql.DB
var rnd *rand.Rand
var savages []savage
