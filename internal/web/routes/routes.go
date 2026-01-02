package routes

import (
	"net/http"
)

var All = map[string]http.HandlerFunc{
	"/":                 HomePage(),
	"GET /site":         SitePage(),
	"POST /site":        CreateSite(),
	"GET /site/{id}":    SitePage(),
	"PUT /site/{id}":    UpdateSite(),
	"DELETE /site/{id}": UpdateSite(),
	"GET /config":       ConfigPage(),
	"PUT /config":       UpdateSite(),
}
