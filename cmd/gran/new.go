package main

import (
	"log"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type subCommandNew struct{}

func (sc *subCommandNew) name() string {
	return "new"
}

func (sc *subCommandNew) usage() string {
	return "create a new Granary secret file"
}

func (sc *subCommandNew) handle(c *cli.Context) error {
	path := c.String("path")
	if util.IsFileExists(path) {
		log.Fatal("file already exists: ", path)
	}

	password, err := util.AskPassword("Enter a new passkey")
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.New(path, password); err != nil {
		log.Fatal(err)
	}

	return nil
}
