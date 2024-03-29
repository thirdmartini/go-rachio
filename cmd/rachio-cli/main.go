package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	appName = "rachio-cli"
)

func commandNotFound(c *cli.Context, command string) {
	log.Fatalf("'%s' is not a %s command. See '%s --help'.",
		command, c.App.Name, c.App.Name)
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Author = "Sebastian Sobolewski"
	app.Email = "spsobole@readybydawn.com"
	app.Usage = "CLI For Rachio Sprinkler Controllers"
	app.Version = "10"

	app.Flags = globalFlags
	app.Commands = commands
	app.CommandNotFound = commandNotFound

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
