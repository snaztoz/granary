package main

import (
	"io"
	"log"
	"os"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type subCommandSet struct{}

func (sc *subCommandSet) name() string {
	return "set"
}

func (sc *subCommandSet) usage() string {
	return "set a secret value"
}

func (sc *subCommandSet) handle(c *cli.Context) error {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 2 {
		log.Fatal("incorrect number of arguments (usage: gran set <KEY> <VALUE>)")
	}

	password, err := util.AskPassword("Enter passkey")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
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

	data[args.Get(0)] = args.Get(1)
	if err := st.WriteData(data); err != nil {
		log.Fatal(err)
	}

	f.Seek(0, io.SeekStart)
	if err := st.Persist(f); err != nil {
		log.Fatal(err)
	}

	return nil
}
