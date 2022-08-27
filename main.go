package main

import (
	"conceitoExato/db"
	"conceitoExato/db/seed"
	"conceitoExato/env"
	"conceitoExato/router"
	"conceitoExato/server"
	"fmt"
)

func main() {
	env.Load()
	db.StartDatabase()
	seed.SeedDatabase()
	currentServer := server.CreateServer()
	router.Avaible(currentServer.GetServerEngine())
	currentServer.GetServerEngine().Run(fmt.Sprintf("%s:%s", env.Server.HOST, env.Server.PORT))
}
