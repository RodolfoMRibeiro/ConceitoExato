package main

import (
	"conceitoExato/adapter/db"
	"conceitoExato/adapter/db/seed"
	"conceitoExato/adapter/env"
	"conceitoExato/adapter/router"
	"conceitoExato/adapter/server"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	env.Load()
	db.StartDatabase()
	seed.SeedDatabase()
	servidor := server.CreateServer()
	router.Avaible(servidor.GetServerEngine())
	servidor.Run()
}
