package main

import (
	"log"

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
	path := c.String("file")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 2 {
		log.Fatal("incorrect number of arguments\n\nusage: gran set <KEY> <VALUE>")
	}

	password, err := util.AskPassword("Enter passkey")
	if err != nil {
		log.Fatal(err)
	}

	st, err := storage.Open(path, password)
	if err != nil {
		log.Fatal(err)
	}

	data, err := st.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	data[args.Get(0)] = args.Get(1)
	if err := st.WriteFile(data); err != nil {
		log.Fatal(err)
	}

	return nil
}
