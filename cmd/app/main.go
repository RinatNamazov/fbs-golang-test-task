package main

import (
	"flag"
	"log"

	"github.com/RinatNamazov/fbs-golang-test-task/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	app, err := app.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
