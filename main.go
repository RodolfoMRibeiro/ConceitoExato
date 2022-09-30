package main

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/seed"
	"conceitoExato/adapter/env"
	"conceitoExato/adapter/router"
	"conceitoExato/adapter/server"
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
