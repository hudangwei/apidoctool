package main

import (
	"log"

	"github.com/urfave/cli"
)

func main() {

	app := NewConfigurator()
	app.App().Action = func(c *cli.Context) error {
		conf := app.Get()
		GenerateMarkDownFile(conf.FilePath)
		return nil
	}
	if err := app.Run(); err != nil {
		log.Fatalf("error on run app, %v", err)
	}
}
