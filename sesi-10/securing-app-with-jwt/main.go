package main

import (
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/database"
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8082")
}