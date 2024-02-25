package commands

import (
	"fmt"
	"log"
	"os"

	"blinders/packages/auth"
	authutils "blinders/packages/auth/utils"
	"blinders/packages/utils"

	firebaseAuth "firebase.google.com/go/auth"
	"github.com/urfave/cli/v2"
)

var client *firebaseAuth.Client

var AuthCommand = cli.Command{
	Name:        "auth",
	Subcommands: []*cli.Command{&loadAuthCommand, &genWSCatCommand},
	Before: func(ctx *cli.Context) error {
		env := ctx.String("env")
		adminJSON, _ := utils.GetFile(fmt.Sprintf("firebase.admin.%v.json", env))
		a, err := auth.NewFirebaseManager(adminJSON)
		if err != nil {
			return err
		}
		client = a.Client

		return nil
	},
}

var loadAuthCommand = cli.Command{
	Name:        "load-user",
	Description: "load user jwt by using uid",
	Args:        true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "uid",
		},
	},
	Action: func(ctx *cli.Context) error {
		webAPIKey := os.Getenv("WEB_API_KEY")
		if webAPIKey == "" {
			log.Fatal("WEB_API_KEY is required from environment")
		}
		env := ctx.String("env")
		var uid string
		if uid = ctx.String("uid"); uid == "" {
			uid = os.Getenv("USER_UID")
		}
		if uid == "" {
			log.Fatal("USER_UID is required from environment")
		}

		cacheFile := fmt.Sprintf("auth.%v.json", env)
		idToken, authToken, err := authutils.LoadFirebaseAuthForUserWithCache(
			client,
			uid,
			webAPIKey,
			cacheFile,
		)
		if err != nil {
			return err
		}

		fmt.Printf("JWT of %v: %v\n", authToken.Firebase.Identities["email"], idToken)

		return nil
	},
}

var genWSCatCommand = cli.Command{
	Name:        "gen-wscat",
	Description: "load user jwt and generate wscat command for easily connect to websocket api by using uid",
	Args:        true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "uid",
		},
		&cli.StringFlag{
			Name:     "endpoint",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		webAPIKey := os.Getenv("WEB_API_KEY")
		if webAPIKey == "" {
			log.Fatal("WEB_API_KEY is required from environment")
		}
		env := ctx.String("env")
		var uid string
		if uid = ctx.String("uid"); uid == "" {
			uid = os.Getenv("USER_UID")
		}
		if uid == "" {
			log.Fatal("USER_UID is required from environment")
		}

		endpoint := ctx.String("endpoint")

		cacheFile := fmt.Sprintf("auth.%v.json", env)
		idToken, _, err := authutils.LoadFirebaseAuthForUserWithCache(
			client,
			uid,
			webAPIKey,
			cacheFile,
		)
		if err != nil {
			return err
		}

		fmt.Printf("wscat -H \"Authorization: Bearer %v\" -c wss://%v\n", idToken, endpoint)

		return nil
	},
}
