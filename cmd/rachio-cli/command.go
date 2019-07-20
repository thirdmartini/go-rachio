package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli"

	"git.thirdmartini.com/pub/go-rachio"
)

var (
	formatFlag = cli.StringFlag{
		Name:  "format",
		Value: "human",
		Usage: "Output display format (json|human|",
	}

	tokenFlag = cli.StringFlag{
		Name:  "token",
		Value: "",
		Usage: "Rachio API Token",
	}

	userIdFlag = cli.StringFlag{
		Name:  "user.id",
		Value: "",
		Usage: "User ID of user to display",
	}
)

var globalFlags = []cli.Flag{
	tokenFlag,
	formatFlag,
}

// Commands are the top-level commands for metastorctl.
var commands = []cli.Command{
	{
		Name:        "device",
		ShortName:   "dev",
		Usage:       "device functions",
		Description: "Device functions",
		Subcommands: deviceCommands,
	},
	{
		Name:        "person",
		Usage:       "user information",
		Description: "Shows user information",
		Flags: []cli.Flag{
			userIdFlag,
		},
		Action: personShow,
	},
}

func mustGetRachio(ctx *cli.Context) *rachio.Rachio {

	token := ctx.GlobalString(tokenFlag.Name)
	if token == "" {
		//token, err := readToken()
	}

	if token == "" {
		log.Fatal("Rachio API token not provided")
	}

	r := rachio.NewClient(token)

	_, err := r.Self()
	if err != nil {
		log.Fatalf("Could not connect to Rachio API: %s\n", err.Error())
	}

	return r
}

func mustGetDisplayFormat(ctx *cli.Context) string {
	switch ctx.GlobalString(formatFlag.Name) {
	case "json":
		return "json"
	case "human":
		return "human"
	}

	return "human"
}

func prettyPrintJSON(obj interface{}) {
	marsh := json.NewEncoder(os.Stdout)
	marsh.SetIndent("", "\t")
	marsh.Encode(obj)
}
