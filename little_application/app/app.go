package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app := cli.NewApp()

	app.Name = "App Search IP"
	app.Usage = "Found IPS"
	app.Commands = []cli.Command{

		{
			Name:   "ip",
			Usage:  "Busca ips de enderecos",
			Flags:  flags,
			Action: buscarIP,
		},
		{
			Name:   "server",
			Usage:  "Busca servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

func buscarIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {

	server := c.String("host")

	serv, err := net.LookupNS(server)

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range serv {
		fmt.Println(s.Host)
	}

}
