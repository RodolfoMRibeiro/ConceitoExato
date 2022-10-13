package main

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/seed"
	"conceitoExato/adapter/env"
	"conceitoExato/adapter/router"
	"conceitoExato/adapter/server"
)

func main() {
	env.Load()
	db.StartDatabase()
	seed.SeedDatabase()
	servidor := server.CreateServer()
	router.Avaible(servidor.GetServerEngine())
	servidor.Run()
}
