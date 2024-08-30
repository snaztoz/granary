package main

import (
	"fmt"
	"log"

	"github.com/snaztoz/granary/internal/storage"
	"github.com/snaztoz/granary/internal/util"
	"github.com/urfave/cli/v2"
)

type subCommandGet struct{}

func (sc *subCommandGet) name() string {
	return "get"
}

func (sc *subCommandGet) usage() string {
	return "get secret value"
}

func (sc *subCommandGet) handle(c *cli.Context) error {
	path := c.String("file")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}

	args := c.Args()
	if args.Len() != 1 {
		log.Fatal("incorrect number of arguments (usage: gran get <KEY>)")
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

	value, exist := data[c.Args().First()]
	if exist {
		fmt.Println(value)
	}

	return nil
}
