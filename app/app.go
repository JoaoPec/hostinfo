package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "hostinfo"
	app.Usage = "Search ip's and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{

		{
			Name:   "ip",
			Usage:  "Search IP'S from adresses in the web",
			Flags:  flags,
			Action: searchIps,
		},
        {
            Name: "server",
            Usage: "Search server names from adresses in the web",
            Flags: flags,
            Action: searchServer,
        },
	}

	return app
}

func searchServer(c *cli.Context){

    host := c.String("host")

    servers,err := net.LookupNS(host) // name server
    if err != nil{
        log.Fatal(err)
    }

    for _,server := range servers{
        fmt.Println(server.Host)
    }

}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host) //Ip
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
