package main

import (
	"github.com/codegangsta/cli"
	"github.com/ShaneKilkelly/jetcan/localstorage"
	"github.com/ShaneKilkelly/jetcan/config"
	"fmt"
	"os"
)

type Jetcan struct {
	Config	*config.Config
	Storage	*localstorage.LocalStorage
}

func NewJetcan() (*Jetcan, error) {
	var (
		cfg	*config.Config
		l	*localstorage.LocalStorage
		j	*Jetcan
		err	error
	)

	j = &Jetcan{}

	cfg, err = config.Load()
	if err != nil {
		return nil, err
	} else {
		j.Config = cfg
	}

	l, err = localstorage.New(j.Config)
	if err != nil {
		return nil, err
	} else {
		j.Storage = l
	}

	return j, nil
}

func handler(c *cli.Context) {
	jetcan, err := NewJetcan()
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	fmt.Println("\nthis is jetcan", jetcan.Config, jetcan.Storage, "\n")
	fmt.Println(jetcan.Storage.GetAuthToken())
}

func initCliApp() (app *cli.App, err error) {

	app = cli.NewApp()

	app.Name = "jetcan"
	app.Version = "0.1.0"
	app.Action = handler

	return app, nil
}

func main() {
	app , err := initCliApp()
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		app.Run(os.Args)
	}
}
