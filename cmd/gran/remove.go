package main

import (
	"log"

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
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 1 {
		log.Fatal("incorrect number of arguments (usage: gran remove <KEY>)")
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

	delete(data, c.Args().First())
	if err := st.WriteFile(data); err != nil {
		log.Fatal(err)
	}

	return nil
}
