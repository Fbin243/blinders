package main

import (
	"fmt"
	"log"
	"os"

	"blinders/packages/auth"
	"blinders/packages/db"
	"blinders/packages/utils"
	restapi "blinders/services/rest/api"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var apiManager restapi.Manager

func init() {
	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatal("failed to load env", err)
	}

	app := fiber.New()
	adminJSON, _ := utils.GetFile("firebase.admin.development.json")
	auth, _ := auth.NewFirebaseManager(adminJSON)

	url := fmt.Sprintf(
		db.MongoURLTemplate,
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
	)
	db := db.NewMongoManager(url, os.Getenv("MONGO_DATABASE"))
	if db == nil {
		log.Fatal("cannot create database manager")
	}

	apiManager = *restapi.NewManager(app, auth, db)
	_ = apiManager.InitRoute(restapi.InitOptions{})
}

func main() {
	port := os.Getenv("REST_API_PORT")
	err := apiManager.App.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Println("launch chat service error", err)
	}
}
