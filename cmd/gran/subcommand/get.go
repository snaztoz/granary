package subcommand

import (
	"fmt"
	"log"
	"os"

	"github.com/snaztoz/granary/cmd/gran/util"
	"github.com/snaztoz/granary/internal/storage"
	"github.com/urfave/cli/v2"
)

type Get struct{}

func (sc *Get) Name() string {
	return "get"
}

func (sc *Get) Usage() string {
	return "get a secret value"
}

func (sc *Get) Handle(c *cli.Context) error {
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

	value, exist := data[c.Args().First()]
	if exist {
		fmt.Println(value)
	}

	return nil
}

func (sc *Get) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (sc *Get) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 1 {
		log.Fatal("incorrect number of arguments (usage: gran get <KEY>)")
	}
}
