package main

import (
	"os"

	"github.com/gen1us2k/log"
	"github.com/urfave/cli"
)

var (
	Version  string
	LogLevel string
	FilePath string
)

type Configuration struct {
	customConfig *CustomConfig
	app          *cli.App
}

func NewConfigurator() *Configuration {
	Version = "0.0.0"
	app := cli.NewApp()
	app.Name = "apidoctool"
	app.Usage = "generate markdown docs with swagger spec"
	return &Configuration{
		customConfig: &CustomConfig{},
		app:          app,
	}
}

func (c *Configuration) fillConfig() *CustomConfig {
	return &CustomConfig{
		FilePath: FilePath,
	}
}
func (c *Configuration) After(cb func(c *cli.Context) error) {
	c.app.After = func(ctx *cli.Context) error {
		log.SetLevel(log.MustParseLevel(LogLevel))
		c.customConfig = c.fillConfig()
		return cb(ctx)
	}
}

func (c *Configuration) Run() error {
	c.app.Before = func(ctx *cli.Context) error {
		log.SetLevel(log.MustParseLevel(LogLevel))
		return nil
	}
	c.app.Flags = c.setupFlags()
	return c.app.Run(os.Args)
}

func (c *Configuration) App() *cli.App {
	return c.app
}

func (c *Configuration) setupFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "loglevel, l",
			Value:       "debug",
			Usage:       "set log level",
			Destination: &LogLevel,
			EnvVar:      "LOG_LEVEL",
		},
		cli.StringFlag{
			Name:        "FilePath, f",
			Value:       "swagger.json",
			Usage:       "swagger spec file or swagger spec url",
			EnvVar:      "SWAGGER_SPEC_FILEPATH",
			Destination: &FilePath,
		},
	}
}

func (c *Configuration) Get() *CustomConfig {
	c.customConfig = c.fillConfig()
	return c.customConfig
}
