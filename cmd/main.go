package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"

	p1 "gopinger/pinger"
)

/*

pingling

*/

func main() {

	app := &cli.App{
		Name:  "pinger",
		Usage: "An simple application to allow you search for IP's connected on your machine",
		Action: func(ctx *cli.Context) error {
			input := ctx.Args().First()
			if input == "" {
				return fmt.Errorf("please provide an ip address")
			}

			if strings.ToLower(input) == "search" {
				err := p1.PingSearch(p1.GetLocalIP())
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
