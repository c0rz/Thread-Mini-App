package main

import (
	"Blog/config"
	"Blog/routes"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	DB  *sql.DB
)


func main() {
	DB =  config.Database()

	config.DBMigrate(DB, 0)

	defer DB.Close()
	
	gin.SetMode(gin.DebugMode)

	routes := routes.Routes()
	routes.Run(":"+ os.Getenv("PORT"))
}