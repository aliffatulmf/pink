package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Command struct {
	Host        string
	Port        uint
	Environment string
}

func (c *Command) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

const (
	DevelopmentMode = "development"
	ProductionMode  = "production"
)

func ParseServerFlag() *Command {
	var c Command

	f := flag.NewFlagSet("Pink Server", flag.ExitOnError)
	f.StringVar(&c.Host, "host", "localhost", "Host to listen on")
	f.UintVar(&c.Port, "port", 8080, "Port to listen on")
	f.StringVar(&c.Environment, "env", DevelopmentMode, "Environment to run on")
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	return &c
}

var Args = ParseServerFlag()
