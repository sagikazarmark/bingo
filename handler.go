package bingo

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/julienschmidt/httprouter"
)

// Creates the HTTP Handler
func NewHandler(b *Bin) http.Handler {
	router := httprouter.New()

	// Add global endpoints
	for _, endpoint := range b.Endpoints {
		router.Handle(endpoint.Method, endpoint.Uri.Path, endpoint.Handler)
	}

	// Add groupped endpoints
	for _, group := range b.Groups {
		for _, endpoint := range group.Endpoints {
			router.Handle(endpoint.Method, endpoint.Uri.Path, endpoint.Handler)
		}
	}

	// Add index page
	router.GET("/", b.Index)

	// Add assets necessary for the index page
	// Note: this might conflict with any endpoint starting with /_assets
	assets := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "data/assets"}
	router.Handler("GET", "/_assets/*file", http.StripPrefix("/_assets/", http.FileServer(assets)))

	return router
}
