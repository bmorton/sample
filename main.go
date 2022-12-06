package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// This is the example application from https://cli.urfave.org/v2/getting-started/
func main() {
	app := &cli.App{
		Name:  "sample",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say! v1!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
