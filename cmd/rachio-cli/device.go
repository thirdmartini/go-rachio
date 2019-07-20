package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

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
	{
		Name:        "events",
		ShortName:   "ev",
		Usage:       "--device.id=<device>",
		Description: "Shows Device events",
		Action:      deviceShowEvents,
		Flags: []cli.Flag{
			deviceIdFlag,
		},
	},
}

func deviceList(ctx *cli.Context) {
	r := mustGetRachio(ctx)

	self, err := r.Self()
	if err != nil {
		log.Fatal(err)
	}

	switch mustGetDisplayFormat(ctx) {
	case FormatHuman:
		w := new(tabwriter.Writer)
		// Format in tab-separated columns with a tab stop of 8.
		w.Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(w, "ID\tName\tStatus\tModel\tSerial Number\tMAC Addess\t\n")

		for _, device := range self.Devices {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t\n",
				device.ID,
				device.Name,
				device.Status,
				device.Model,
				device.SerialNumber,
				device.MacAddress,
			)
		}
		w.Flush()

	case FormatJSON:
		prettyPrintJSON(self.Devices)
	}
}

func deviceShowEvents(ctx *cli.Context) {
	r := mustGetRachio(ctx)

	device, err := r.Device(ctx.String(deviceIdFlag.Name))
	if err != nil {
		log.Fatal(err)
	}

	time.Now()

	events, err := device.Events(time.Now().AddDate(0, 0, -14), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	switch mustGetDisplayFormat(ctx) {
	case FormatHuman:
		w := new(tabwriter.Writer)
		// Format in tab-separated columns with a tab stop of 8.
		w.Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(w, "Date\tType\tSubtype\tTopic\tSummary\t\n")

		for _, event := range events {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n",
				event.Timestamp().String(),
				event.Type,
				event.SubType,
				event.Topic,
				event.Summary,
			)
		}
		w.Flush()

	case FormatJSON:
		prettyPrintJSON(events)
	}

}
