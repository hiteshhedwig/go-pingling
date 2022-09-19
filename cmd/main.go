package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"

	p1 "gopinger/pinger"
)

func main() {

	app := &cli.App{
		Name:  "pinger",
		Usage: "An simple application to allow you search for IP's connected on your machine",
		Action: func(ctx *cli.Context) error {
			err := p1.PingSearch(ctx.Args().First())
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
