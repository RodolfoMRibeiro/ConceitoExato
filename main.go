package main

import (
	"conceitoExato/db"
	"conceitoExato/env"
	"conceitoExato/router"
	"conceitoExato/server"
	"fmt"
)

func main() {
	env.Load()
	db.StartDatabase()
	currentServer := server.CreateServer()
	router.Avaible(currentServer.GetServerEngine())
	currentServer.GetServerEngine().Run(fmt.Sprintf("%s:%s", env.Server.HOST, env.Server.PORT))
}
