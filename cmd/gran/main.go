package main

import (
	"log"
	"os"

	"github.com/snaztoz/granary/cmd/gran/subcommand"
	"github.com/urfave/cli/v2"
)

type subCommand interface {
	Name() string
	Usage() string
	Handle(c *cli.Context) error
	Flags() []cli.Flag
}

var (
	cliSubCommands = []*cli.Command{}
)

func init() {
	for _, v := range []subCommand{
		&subcommand.Get{},
		&subcommand.List{},
		&subcommand.New{},
		&subcommand.Remove{},
		&subcommand.Set{},
	} {
		cliSubCommands = append(cliSubCommands, &cli.Command{
			Name:   v.Name(),
			Usage:  v.Usage(),
			Action: v.Handle,
			Flags:  v.Flags(),
		})
	}
}

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
				Value:   "secrets",
				Usage:   "Path to Granary secret file",
			},
		},
		Commands: cliSubCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
