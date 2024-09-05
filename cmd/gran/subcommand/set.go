package subcommand

import (
	"io"
	"log"
	"os"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type Set struct{}

func (sc *Set) Name() string {
	return "set"
}

func (sc *Set) Usage() string {
	return "set a secret value"
}

func (sc *Set) Handle(c *cli.Context) error {
	sc.validate(c)

	password, err := util.AskPassword("Enter passkey")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(c.String("path"), os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	st, err := storage.Open(f, password)
	if err != nil {
		log.Fatal(err)
	}

	data, err := st.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	data[c.Args().Get(0)] = c.Args().Get(1)

	if err := st.WriteData(data); err != nil {
		log.Fatal(err)
	}

	f.Seek(0, io.SeekStart)
	if err := st.Persist(f); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (sc *Set) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 2 {
		log.Fatal("incorrect number of arguments (usage: gran set <KEY> <VALUE>)")
	}
}
