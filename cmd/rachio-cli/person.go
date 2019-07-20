package main

import (
	"fmt"
	"log"

	"git.thirdmartini.com/pub/go-rachio"

	"github.com/urfave/cli"
)

func printPerson(info *rachio.Person) {
	fmt.Printf("       ID: %s\n", info.ID)
	fmt.Printf("User Name: %s\n", info.Username)
	fmt.Printf("Full Name: %s\n", info.FullName)
	fmt.Printf("   E-Mail: %s\n", info.Email)
}

func personShow(ctx *cli.Context) {
	r := mustGetRachio(ctx)

	id := ctx.String(userIdFlag.Name)

	var info *rachio.Person
	var err error

	if id == "" {
		info, err = r.Self()
	} else {
		info, err = r.Person(id)
	}

	if err != nil {
		log.Fatal(err)
	}

	switch mustGetDisplayFormat(ctx) {
	case "human":
		printPerson(info)

	case "json":
		prettyPrintJSON(info)
	}
}
