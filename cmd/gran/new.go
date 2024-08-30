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
	path := c.String("file")
	if util.IsFileExists(path) {
		log.Fatalf("file already exists: %s\n\n", path)
	}

	password, err := util.AskPassword("Enter a new passkey")
	if err != nil {
		log.Fatalln(err)
	}

	storage.New(path, password)

	return nil
}
