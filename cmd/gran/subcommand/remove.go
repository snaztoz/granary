package subcommand

import (
	"io"
	"log"
	"os"

	"github.com/snaztoz/granary/cmd/gran/util"
	"github.com/snaztoz/granary/internal/storage"
	"github.com/urfave/cli/v2"
)

type Remove struct{}

func (sc *Remove) Name() string {
	return "remove"
}

func (sc *Remove) Usage() string {
	return "remove a secret"
}

func (sc *Remove) Handle(c *cli.Context) error {
	sc.validate(c)
	path := c.String("path")

	passphrase, err := util.GetPassphrase(path+".gpass", "Enter passphrase")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
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

	delete(data, c.Args().First())
	if err := st.WriteData(data); err != nil {
		log.Fatal(err)
	}

	// remove all previous file content as they will mess up with
	// the new Base64 encoded data
	f.Truncate(0)
	f.Seek(0, io.SeekStart)
	if err := st.Persist(f); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (sc *Remove) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (sc *Remove) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 1 {
		log.Fatal("incorrect number of arguments (usage: gran remove <KEY>)")
	}
}
