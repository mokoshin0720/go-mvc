// controllerをroutingに登録する。server.goから呼び出される
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
	"tutorial/mvc/pkg/controller"
)

func HandleRequests() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Get("/index", controller.Index)
	r.Get("/users", controller.AllUsers)
	r.Get("/userid", controller.Userid)
	r.Post("/create", controller.Create)

	return r
}