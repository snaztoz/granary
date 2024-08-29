package main

import (
	"fmt"
	"log"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type subCommandList struct{}

func (sc *subCommandList) name() string {
	return "list"
}

func (sc *subCommandList) usage() string {
	return "get the list of all secrets inside a Granary secret file"
}

func (sc *subCommandList) handle(c *cli.Context) error {
	path := c.String("file")
	if !util.IsFileExists(path) {
		log.Fatalf("file is not exist: %s\n\n", path)
	}

	password, err := util.AskPassword()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := storage.ReadFile(path, password)
	if err != nil {
		log.Fatalln(err)
	}

	for k := range data {
		fmt.Println(k)
	}

	return nil
}
