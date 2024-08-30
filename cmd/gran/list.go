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
	return "get the list of all stored secrets"
}

func (sc *subCommandList) handle(c *cli.Context) error {
	path := c.String("file")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
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

	fmt.Println(data)

	return nil
}
