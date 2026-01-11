package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/p-tupe/pingmon/internal/web"
)

type Routes map[string]http.HandlerFunc

func InitRoutes(jobs []*Ping) Routes {
	tmpl, err := template.ParseGlob("./internal/web/public/*.html")
	if err != nil {
		log.Fatalf("Error while parsing templates, %v", err)
	}

	var routes = Routes{
		"GET /{$}": web.HomePage(tmpl, jobs),

		"GET /site":         web.SitePage(tmpl),
		"POST /site":        web.CreateSite(),
		"GET /site/{id}":    web.SitePage(tmpl),
		"PUT /site/{id}":    web.UpdateSite(),
		"DELETE /site/{id}": web.UpdateSite(),

		"GET /config": web.ConfigPage(tmpl),
		"PUT /config": web.UpdateSite(),

		"GET /static/{asset}": func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./internal/public/static/"+r.PathValue("asset"))
		},
	}

	return routes
}
