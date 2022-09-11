package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CommandLine"
	app.Usage = "Find IP's and others informations"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find IP's",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "example.com.br",
				},
			},
			Action: func(c *cli.Context) error {
				FindIP(c)
				return nil
			},
		},
	}

	return app
}

func FindIP(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
