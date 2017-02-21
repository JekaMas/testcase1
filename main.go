package main

import (
	"flag"
	"fmt"

	"generator/app"
	"generator/app/config"

	"strconv"
)

var port string

func main() {
	flag.StringVar(&port, "port", "", "Applicatoin port")
	flag.Parse()

	conf := config.New()
	portValue, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while starting server. Incorrect port: %v", port))
	}

	conf.Port = portValue
	application := app.NewApp(conf)

	err = application.Bootstrap()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while starting server: %v", err))
	}

	application.Run()
}
