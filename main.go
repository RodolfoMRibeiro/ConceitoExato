package main

import (
	"conceitoExato/db"
	"conceitoExato/env"
)

func main() {
	env.Load()
	db.StartDatabase()
}
