package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/luytbq/astrio-authentication-service/cmd/api"
	"github.com/luytbq/astrio-authentication-service/config"
	"github.com/luytbq/astrio-authentication-service/internal/database"
)

func main() {
	db, err := database.NewPosgresDB()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(config.App.SERVER_PORT, config.App.SERVER_API_PREFIX, db)

	server.Run()
}
