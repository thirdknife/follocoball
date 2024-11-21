package server

import (
	"follocoball/controller"
	"follocoball/static"
	"follocoball/views"
	"net/http"

	"github.com/go-fuego/fuego"
)

type Resources struct {
	API   controller.Resource
	Views views.PlayerResource
}

func (rs Resources) Setup(
	options ...func(*fuego.Server),
) *fuego.Server {
	serverOptions := []func(*fuego.Server){
		fuego.WithAddr(":9999"),
		fuego.WithGlobalResponseTypes(http.StatusForbidden, "Forbidden"),
	}
	options = append(serverOptions, options...)

	app := fuego.NewServer(options...)
	app.OpenApiSpec.Info.Title = "Gourmet API"

	fuego.Register(app, fuego.Route[any, any]{
		Path:     "/static/",
		Method:   http.MethodGet,
		FullName: "static handler",
	}, http.StripPrefix("/static", static.Handler()), func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age=600")

			h.ServeHTTP(w, r)
		})
	})

	rs.API.MountRoutes(fuego.Group(app, "/api"))
	rs.Views.MountRoutes(fuego.Group(app, "/"))

	return app
}
