package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("sweep", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"": func() (cli.Command, error) {
			return &Sweep{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
