package main

import (
	"io"
	"log"
	"os"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type subCommandRemove struct{}

func (sc *subCommandRemove) name() string {
	return "remove"
}

func (sc *subCommandRemove) usage() string {
	return "remove a secret"
}

func (sc *subCommandRemove) handle(c *cli.Context) error {
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

	delete(data, c.Args().First())
	if err := st.WriteData(data); err != nil {
		log.Fatal(err)
	}

	f.Seek(0, io.SeekStart)
	if err := st.Persist(f); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (sc *subCommandRemove) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 1 {
		log.Fatal("incorrect number of arguments (usage: gran remove <KEY>)")
	}
}
