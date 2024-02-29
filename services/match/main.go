package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"blinders/packages/auth"
	"blinders/packages/db"
	"blinders/packages/match"
	"blinders/packages/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	matchapi "blinders/services/match/api"
)

var service matchapi.Service

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	adminJSON, _ := utils.GetFile("firebase.admin.json")
	authManager, err := auth.NewFirebaseManager(adminJSON)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	url := fmt.Sprintf(
		db.MongoURLTemplate,
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
	)

	db := db.NewMongoManager(url, os.Getenv("MONGO_DATABASE"))

	fmt.Println("Connect to mongo url", url)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	core := match.NewMongoMatcher(db, redisClient, match.MockEmbedder{})

	service = matchapi.Service{Auth: authManager, App: app, Core: core}
	service.InitRoute()
}

func main() {
	port := os.Getenv("MATCH_SERVICE_PORT")
	log.Panic(service.App.Listen(":" + port))
}
