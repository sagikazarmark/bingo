package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	flag "github.com/spf13/pflag"
)

var port string

func init() {
    flag.StringVar(&port, "port", "8080", "Custom port")
}

func main() {
	bin := New("bingo", "Bingo example application", "Lorem ipsum")
	bin.Endpoints = []*Endpoint{
		&Endpoint {
			Method: "GET",
			Path: "/endpoint",
			Description: "Lorem ipsum dolor",
		},
		&Endpoint {
			Method: "GET",
			Path: "/endpoint/:param",
			Description: "Lorem ipsum dolor",
			Defaults: map[string]string{
				"param": "100",
			},
			Handler: func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				fmt.Fprintf(w, "You wrote: %s", p.ByName("param"))
			},
		},
		&Endpoint {
			Method: "POST",
			Path: "/endpoint/post",
			Description: "...sit amet",
		},
	}

	bin.Groups = []*Group{
		&Group {
			Name: "Group",
			Description: "Lorem ipsum dolor",
			Endpoints: []*Endpoint{
				&Endpoint {
					Method: "GET",
					Path: "/group/endpoint",
					Description: "Lorem ipsum dolor",
				},
				&Endpoint {
					Method: "GET",
					Path: "/group/endpoint/:param",
					Description: "Lorem ipsum dolor",
					Defaults: map[string]string{
						"param": "100",
					},
					Handler: func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						fmt.Fprintf(w, "You wrote: %s", p.ByName("param"))
					},
				},
				&Endpoint {
					Method: "POST",
					Path: "/group/endpoint/post",
					Description: "...sit amet",
				},
			},
		},
	}

	log.Println("Starting server on *:" + port)
	log.Fatal(http.ListenAndServe(":" + port, CreateHandler(bin)))
}
