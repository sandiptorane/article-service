package main

import (
	"article-service/internal/database/connection"
	"article-service/internal/handler"
	"article-service/internal/router"
	"article-service/internal/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env config
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal("not able to load env", err)
	}

	// get db connection
	conn, err := connection.GetConnection()
	if err != nil {
		log.Fatal("db connection err:", err)
		return
	}

	// get database instance
	dbInstance := service.GetInstance(conn)

	// get new handler
	app := handler.New(dbInstance)

	// register routes
	r := router.RegisterRoutes(app)

	// run server
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("run error:", err)
	}
}
