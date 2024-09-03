package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	scGet    = subCommandGet{}
	scList   = subCommandList{}
	scNew    = subCommandNew{}
	scRemove = subCommandRemove{}
	scSet    = subCommandSet{}
)

func main() {
	// Remove timestamp
	log.SetFlags(0)

	app := cli.App{
		Name:  "gran",
		Usage: "store your secrets inside an encrypted file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   "secrets.gran",
				Usage:   "Path to Granary secret file",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   scGet.name(),
				Usage:  scGet.usage(),
				Action: scGet.handle,
			},
			{
				Name:   scList.name(),
				Usage:  scList.usage(),
				Action: scList.handle,
			},
			{
				Name:   scNew.name(),
				Usage:  scNew.usage(),
				Action: scNew.handle,
			},
			{
				Name:   scRemove.name(),
				Usage:  scRemove.usage(),
				Action: scRemove.handle,
			},
			{
				Name:   scSet.name(),
				Usage:  scSet.usage(),
				Action: scSet.handle,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
