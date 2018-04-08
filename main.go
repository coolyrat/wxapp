package main

import (
	"fmt"
	"github.com/coolyrat/wxapp/app"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"os"
)

func main() {

	var config app.Config

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug app",
			Destination: &config.Debug,
		},
		cli.IntFlag{
			Name:        "server-port",
			Usage:       "server port",
			Value:       8080,
			Destination: &config.Port,
		},
	}

	app.Action = func(c *cli.Context) error {
		initLog(config)
		initRouter(config)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func initRouter(cfg app.Config) {
	root := mux.NewRouter()

	get := root.Methods("GET").Subrouter()

	get.HandleFunc("/hello", hello)
	log.Info("start server...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), root)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}

func initLog(cfg app.Config) {
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

}
