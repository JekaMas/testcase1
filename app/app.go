package app

import (
	"net/http"
	"strconv"

	"context"
	"generator/app/config"
	"generator/app/inject"
	"generator/app/storage"
	"log"
)

//App - application component
type App struct {
	Config   *config.Config
	Router   http.Handler
	DB       storage.CampaignsDAO
	Ctx      context.Context
	TypedCtx inject.FullCtx
}

//Application - application interface for run bootstrap and stop
type Application interface {
	Run() error
	Bootstrap() error
}

//NewApp create new application component
func NewApp(conf *config.Config) *App {
	return &App{
		Config: conf,
	}
}

//Run running all http servers
func (this *App) Run() error {
	http.Handle("/", this.Router)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(this.Config.GetPort()),
		Handler: nil,
	}

	log.Print("Campaign service started on port ", this.Config.GetPort())

	return server.ListenAndServe()

}
