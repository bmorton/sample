package main

import (
	"errors"
	"fmt"
	"github.com/creativeprojects/go-selfupdate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

const version = "1.0.2"

// This is the example application from https://cli.urfave.org/v2/getting-started/
func main() {
	app := &cli.App{
		Name:  "sample",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Printf("boom! I say! v%s!\n", version)
			return nil
		},
		Commands: []*cli.Command{updateCommand},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var updateCommand = &cli.Command{
	Name:  "update",
	Usage: "Update the application to the latest version",
	Action: func(c *cli.Context) error {
		return update(version)
	},
}

func update(version string) error {
	latest, found, err := selfupdate.DetectLatest("bmorton/sample")
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %w", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return errors.New("could not locate executable path")
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %w", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}
