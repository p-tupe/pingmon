package web

import (
	"embed"
	"net/http"
)

//go:embed public/*.html
var publicFS embed.FS

//go:embed public/static
var staticFS embed.FS

var Routes = map[string]http.HandlerFunc{
	"GET /{$}": HomePage(),

	"GET /site":         SitePage(),
	"POST /site":        CreateSite(),
	"GET /site/{id}":    SitePage(),
	"PUT /site/{id}":    UpdateSite(),
	"DELETE /site/{id}": UpdateSite(),

	"GET /config": ConfigPage(),
	"PUT /config": UpdateSite(),

	"GET /static/{asset}": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticFS, "public/static/"+r.PathValue("asset"))
	},
}
