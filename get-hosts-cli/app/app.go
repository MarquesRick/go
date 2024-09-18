package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Process() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search IPs and server names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs on internet",
			Flags:  flags,
			Action: SearchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search server name on internet",
			Flags:  flags,
			Action: SearchServer,
		},
	}
	return app
}

func SearchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func SearchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // name server
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
