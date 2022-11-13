package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"

	p1 "gopinger/pinger"

	"golang.org/x/crypto/ssh/terminal"
)

/*

pingling

*/

func main() {

	app := &cli.App{
		Name: "pinger",
		Usage: ` An simple application to allow you search for IP's connected on your machine
		
		search : for list of IP's connected on your machine
			ex:
				pingling search 

		ssh    : SSH into the provided machine
			ex:
				pingling ssh <IP> <Port>
		`,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "22",
				Usage: "ssh port",
			},
		},
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

			if strings.ToLower(input) == "ssh" {
				ip := ctx.Args().Get(1)
				fmt.Print("Enter username (remote machine ): ")
				var username string
				_, err := fmt.Scanln(&username)
				if err != nil {
					return fmt.Errorf("aborted : %v", err)
				}

				fmt.Print("Enter password (remote machine ): ")
				password, err := terminal.ReadPassword(0)
				if err != nil {
					return fmt.Errorf("aborted : %v", err)
				}

				port := ctx.String("port")
				hostport := fmt.Sprintf("%s:%s", ip, port)
				err = p1.Sshclient(hostport, username, string(password))
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
