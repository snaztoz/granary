package subcommand

import (
	"fmt"
	"log"
	"os"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type List struct{}

func (sc *List) Name() string {
	return "list"
}

func (sc *List) Usage() string {
	return "get the list of all stored secrets"
}

func (sc *List) Handle(c *cli.Context) error {
	sc.validate(c)

	passphrase, err := util.AskPassphrase("Enter passphrase")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(c.String("path"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	st, err := storage.Open(f, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	data, err := st.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	if len(data) > 0 {
		fmt.Println(data)
	}

	return nil
}

func (sc *List) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}
}
