package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/appengine"

	"github.com/micnncim/appengine-protobuf-api/src/route"
)

func main() {
	r := chi.NewRouter()
	route.Register(r)
	http.Handle("/", r)
	appengine.Main()
}
