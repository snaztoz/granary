package subcommand

import (
	"log"
	"os"

	"github.com/snaztoz/granary/cmd/gran/util"
	"github.com/snaztoz/granary/internal/storage"
	"github.com/urfave/cli/v2"
)

type New struct{}

func (sc *New) Name() string {
	return "new"
}

func (sc *New) Usage() string {
	return "create a new Granary secret file"
}

func (sc *New) Handle(c *cli.Context) error {
	sc.validate(c)
	path := c.String("path")

	passphrase, err := util.PromptPassphrase("Enter passphrase")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := storage.Init(f, passphrase); err != nil {
		log.Fatal(err)
	}

	if createPassfile := c.Bool("create-passfile"); createPassfile {
		sc.createPassfile(path+".gpass", passphrase)
	}

	return nil
}

func (sc *New) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "create-passfile",
			Value: false,
			Usage: "Create a new passphrase file as well",
		},
	}
}

func (sc *New) validate(c *cli.Context) {
	path := c.String("path")
	if util.IsFileExists(path) {
		log.Fatal("file already exists: ", path)
	}
}

func (sc *New) createPassfile(path, passphrase string) {
	if util.IsFileExists(path) {
		log.Fatal("file already exists: ", path)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(passphrase); err != nil {
		log.Fatal(err)
	}
}
