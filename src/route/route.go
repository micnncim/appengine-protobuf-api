package route

import (
	"github.com/go-chi/chi"

	"github.com/micnncim/appengine-protobuf-api/src/handler"
)

func Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Post("/v1/Echo", handler.Echo)
	})
}
