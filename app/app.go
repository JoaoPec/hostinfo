package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "hostinfo"
	app.Usage = "Search ip's and server names"

	app.Commands = []cli.Command{

		{
			Name:  "ip",
			Usage: "Search IP'S from adresses in the web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
