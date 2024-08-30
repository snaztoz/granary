package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Remove timestamp
	log.SetFlags(0)

	app := cli.App{
		Name:  "gran",
		Usage: "store your secrets inside an encrypted file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "secrets.gran",
				Usage:   "Path to Granary secret file",
			},
		},
		Commands: []*cli.Command{
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
				Name:   scSet.name(),
				Usage:  scSet.usage(),
				Action: scSet.handle,
			},
			// {
			// 	Name:   commands.GetCommandName,
			// 	Usage:  commands.GetCommandUsage,
			// 	Action: commands.HandleGetCommand,
			// },
			// {
			// 	Name:  "set",
			// 	Usage: "set a single key-value pair to Granary secret file",
			// 	Action: func(c *cli.Context) error {
			// 		fmt.Println("setting a single key-value")
			// 		return nil
			// 	},
			// },
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
