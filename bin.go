package main

//go:generate go-bindata data/...

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/elazarl/go-bindata-assetfs"
)

type Endpoint struct {
	Method string
	Path string
	Handler httprouter.Handle
	Description string
	Defaults map[string]string
}

type Group struct {
	Name string
	Description string
	Endpoints []*Endpoint
}

type Bin struct {
	Name string
	ShortDescription string
	Description string
	Endpoints []*Endpoint
	Groups []*Group
	IndexTemplate string
}

// Creates a new Bin
func New(name string, shortDescription string, description string) (*Bin) {
	return &Bin {
		Name: name,
		ShortDescription: shortDescription,
		Description: description,
	}
}

// Creates the HTTP Handler
func CreateHandler(b *Bin) (http.Handler) {
	router := httprouter.New()

	// Add global endpoints
	for _, endpoint := range b.Endpoints {
		router.Handle(endpoint.Method, endpoint.Path, endpoint.Handler)
	}

	// Add groupped endpoints
	for _, group := range b.Groups {
		for _, endpoint := range group.Endpoints {
			router.Handle(endpoint.Method, endpoint.Path, endpoint.Handler)
		}
	}

	// Add index page
	router.GET("/", b.Index)

	// Add assets necessary for the index page
	// Note: this might conflict with any endpoint starting with /_assets
	assets := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "data/assets"}
	router.Handler("GET", "/_assets/*file", http.StripPrefix("/_assets/", http.FileServer(assets)))

	return router;
}

// The index page listing all endpoints
func (b *Bin) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if b.IndexTemplate == "" {
		tplByte, _ := Asset("data/templates/index.html")

		// Avoid multiple template resolution
		b.IndexTemplate = string(tplByte)
	}

	tpl := b.IndexTemplate

	t, _ := template.New("index").Parse(tpl)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	_ = t.Execute(w, b)
}

// Prepare the URL with sane defaults
func (e *Endpoint) PrepareWithDefaults() string {
	path := e.Path

	for k, v := range e.Defaults {
		path = strings.Replace(path, ":" + k, v, -1)
	}

	return path
}
