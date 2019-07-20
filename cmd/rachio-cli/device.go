package main

import (
	"log"

	"github.com/urfave/cli"
)

var deviceCommands = []cli.Command{
	{
		Name:        "list",
		ShortName:   "ls",
		Usage:       "List devices",
		Description: "Lists all devices",
		Action:      deviceList,
		Flags:       []cli.Flag{},
	},
}

func deviceList(ctx *cli.Context) {
	r := mustGetRachio(ctx)

	self, err := r.Self()
	if err != nil {
		log.Fatal(err)
	}

	switch mustGetDisplayFormat(ctx) {
	case "human":

		for _, device := range self.Devices {
			prettyPrintJSON(device)
		}

	case "json":
		prettyPrintJSON(self.Devices)
	}

}
