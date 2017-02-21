package app

import (
	"github.com/gorilla/mux"

	"generator/app/inject"
	"generator/app/transport/gokit"
)

//InitRouter - init router for net/http
func (this *App) InitRouter(r *mux.Router) error {
	this.initServiceFunctions(r, this.TypedCtx)
	this.initRoutes(r, this.TypedCtx)
	this.Router = r
	return nil
}

const (
	URLUser           = "/user"
	URLCampaign       = "/campaign"
	URLImportCampaign = "/import_camp"
	URLSearch         = "/search"
	URLAutoSearch     = "/search_auto"
)

func (this *App) initRoutes(r *mux.Router, ctx inject.FullCtx) {
	r.Handle(URLCampaign, gokit.GetCampaigns(ctx)).Methods("GET")
	r.Handle(URLUser, gokit.GetUser(ctx)).Methods("GET")
	r.Handle(URLImportCampaign, gokit.StoreCampaigns(ctx)).Methods("POST")
	r.Handle(URLSearch, gokit.Search(ctx)).Methods("POST")
	r.Handle(URLAutoSearch, gokit.AutoSearch(ctx)).Methods("GET")
}

func (this *App) initServiceFunctions(r *mux.Router, ctx inject.FullCtx) {
	r.NotFoundHandler = gokit.HandleNotFound(ctx)
}
