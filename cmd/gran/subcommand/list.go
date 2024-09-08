package subcommand

import (
	"fmt"
	"log"
	"os"

	"github.com/snaztoz/granary/cmd/gran/util"
	"github.com/snaztoz/granary/internal/storage"
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
	path := c.String("path")

	passphrase, err := util.GetPassphrase(path+".gpass", "Enter passphrase")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
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
