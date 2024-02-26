package db

import (
	"context"
	"log"
	"time"

	"blinders/packages/db/repo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// username:password@host:port/database
const MongoURLTemplate = "mongodb://%s:%s@%s:%s/%s"

const (
	UserCollection         = "users"
	ConversationCollection = "conversations"
	MessageCollection      = "messages"
)

type MongoManager struct {
	Client        *mongo.Client
	Database      string
	Users         *repo.UsersRepo
	Conversations *repo.ConversationsRepo
	Messages      *repo.MessagesRepo
}

func NewMongoManager(url string, name string) *MongoManager {
	ctx, cal := context.WithTimeout(context.Background(), time.Second*10)
	defer cal()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Println("cannot connect to mongo", err)
		return nil
	}

	return &MongoManager{
		Client:        client,
		Database:      name,
		Users:         repo.NewUsersRepo(client.Database(name).Collection(UserCollection)),
		Conversations: repo.NewConversationsRepo(client.Database(name).Collection(ConversationCollection)),
		Messages:      repo.NewMessagesRepo(client.Database(name).Collection(MessageCollection)),
	}
}
