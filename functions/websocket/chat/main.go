package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	wschat "blinders/functions/websocket/chat/core"
	"blinders/packages/db"
	"blinders/packages/session"
	"blinders/packages/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	agm "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/redis/go-redis/v9"
)

var (
	cfg       aws.Config
	apiClient *agm.Client
)

func init() {
	// TODO: need to store these secrets to aws secret manager instead of pass in env
	sessionManager := session.NewManager(redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}))

	url := fmt.Sprintf(
		db.MongoURLTemplate,
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
	)

	fmt.Println(url)
	database := db.NewMongoManager(url, os.Getenv("MONGO_DATABASE"))
	if database == nil {
		log.Fatal("cannot create database manager")
	}

	wschat.InitApp(sessionManager, database)

	var err error
	cfg, err = config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("failed to load aws config", err)
	}
}

func HandleRequest(
	ctx context.Context,
	req events.APIGatewayWebsocketProxyRequest,
) (any, error) {
	if apiClient == nil {
		apiClient = NewAPIGatewayManagementClient(&cfg, req.RequestContext.DomainName, req.RequestContext.Stage)
	}

	connectionID := req.RequestContext.ConnectionID
	userID := req.RequestContext.Authorizer.(map[string]interface{})["principalId"].(string)

	genericEvent, err := utils.JSONConvert[wschat.ChatEvent](req.Body)
	if err != nil {
		return nil, err
	}

	switch genericEvent.Type {
	case wschat.UserSendMessage:
		payload, err := utils.JSONConvert[wschat.UserSendMessagePayload](req.Body)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "invalid send message event",
			}, nil
		}

		dCh, err := wschat.HandleSendMessage(userID, connectionID, *payload)
		if err != nil {
			log.Println("failed to send message", err)
			return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
		}

		wg := sync.WaitGroup{}
		for {
			d := <-dCh
			if d == nil {
				log.Println("distribute message channel closed")
				break
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				data, err := json.Marshal(d.Payload)
				if err != nil {
					log.Println("can not marshal data", err)
					return
				}

				err = Publish(ctx, d.ConnectionID, data)
				if err != nil {
					log.Println("can not publish message", err)
				}
			}()
		}

		wg.Wait()
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
