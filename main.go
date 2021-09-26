package main

import (
	"net/http"

	"github.com/ngamux/middleware/static"
	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.NewNgamux()
	mux.Use(static.New())

	mux.Get("/", IndexController)
	mux.Get("/blog", BlogController)

	panic(http.ListenAndServe(":8080", mux))
}
