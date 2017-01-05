package bingo

//go:generate go-bindata -pkg bingo data/...

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Bin struct {
	Name             string
	ShortDescription string
	Description      string
	Endpoints        []*Endpoint
	Groups           []*Group
	IndexTemplate    string
}

// Creates a new Bin
func New(name string, shortDescription string, description string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("Name must not be empty")
	}

	return &Bin{
		Name:             name,
		ShortDescription: shortDescription,
		Description:      description,
	}, nil
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

// Idiomatic way of adding an Endpoint without initializing the underlying slice
func (b *Bin) AddEndpoint(endpoint *Endpoint) {
	b.Endpoints = append(b.Endpoints, endpoint)
}

// Idiomatic way of adding a Group without initializing the underlying slice
func (b *Bin) AddGroup(group *Group) {
	b.Groups = append(b.Groups, group)
}
