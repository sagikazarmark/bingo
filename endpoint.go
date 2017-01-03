package bingo

import (
	"errors"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Parameters map[string]string

type Endpoint struct {
	Method string
	Path string
	Handler httprouter.Handle
	Description string
	Parameters Parameters
	Query url.Values
}

func NewEndpoint(method string, path string, handler httprouter.Handle) (*Endpoint, error) {
	if method == "" {
		return nil, errors.New("Method must not be empty")
	}

	if path == "" {
		return nil, errors.New("Path must not be empty")
	}

	if handler == nil {
		return nil, errors.New("Handler must not be empty")
	}

	return &Endpoint {
		Method: method,
		Path: path,
		Handler: handler,
		Parameters: make(Parameters),
		Query: make(url.Values),
	}, nil
}

// Prepare the URL with sane defaults
func (e *Endpoint) Url() string {
	path := e.Path

	for k, v := range e.Parameters {
		path = strings.Replace(path, ":" + k, v, -1)
	}

	if len(e.Query) > 0 {
		path = path + "?" + e.Query.Encode()
	}

	return path
}

func(p Parameters) Set(key string, value string) {
	p[key] = value
}

