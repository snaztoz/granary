package main

import (
	"fmt"
	"log"
	"os"

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
	sc.validate(c)

	password, err := util.AskPassword("Enter passkey")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(c.String("path"))
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

	if len(data) > 0 {
		fmt.Println(data)
	}

	return nil
}

func (sc *subCommandList) validate(c *cli.Context) {
	path := c.String("path")
	if !util.IsFileExists(path) {
		log.Fatal("file is not exist: ", path)
	}
}
