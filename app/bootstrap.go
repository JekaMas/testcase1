package app

import (
	"context"
	"log"

	"github.com/gorilla/mux"

	"generator/app/inject"
	"generator/app/storage/repository/memory"
)

//Bootstrap - init application dependency
func (this *App) Bootstrap() error {
	if err := this.initDatabase(); err != nil {
		return err
	}

	this.TypedCtx = inject.NewTyped(this.DB)

	if err := this.InitRouter(mux.NewRouter()); err != nil {
		log.Print("error init http router")
		return err
	}

	return nil
}

func (this *App) initDatabase() error {
	this.DB = memory.NewStorage()
	this.Ctx = context.WithValue(this.Ctx, inject.DB, this.DB)
	return nil
}
