package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/saltbo/opensae/engine"
)

const (
	FLAG_DSN = "dsn"
)

func main() {
	app := cli.NewApp()
	app.Compiled = time.Now()
	//app.Version = version.Long
	app.Name = "opensae"
	app.Usage = "opensae"
	app.Copyright = "(c) 2020 saltbo.cn"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  FLAG_DSN,
			Usage: "specify data source name",
			Value: "mongodb://root:admin@localhost:27017",
		},
	}
	app.Action = serverAction
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func serverAction(c *cli.Context) {
	ee, err := engine.New(c.String(FLAG_DSN))
	if err != nil {
		log.Fatalln(err)
	}

	if err := ee.Boot(); err != nil {
		log.Fatalln(err)
	}
}
