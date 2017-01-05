package bingo

import (
	"errors"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Parameters map[string]string

type Endpoint struct {
	Method      string
	Uri         *url.URL
	Handler     httprouter.Handle
	Description string
	Parameters  Parameters
	Query       url.Values
}

type EndpointCollector interface {
	AddEndpoint(*Endpoint)
}

func NewEndpoint(method string, uri string, handler httprouter.Handle) (*Endpoint, error) {
	if method == "" {
		return nil, errors.New("Method must not be empty")
	}

	if uri == "" {
		return nil, errors.New("URI must not be empty")
	}

	if handler == nil {
		return nil, errors.New("Handler must not be empty")
	}

	parsedUri, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return &Endpoint{
		Method:     method,
		Uri:        parsedUri,
		Handler:    handler,
		Parameters: make(Parameters),
		Query:      make(url.Values),
	}, nil
}

// Prepare the URI with sane defaults
func (e *Endpoint) BuildUri() string {
	uri := e.Uri.Path

	for k, v := range e.Parameters {
		uri = strings.Replace(uri, ":"+k, v, -1)
	}

	if len(e.Query) > 0 {
		uri = uri + "?" + e.Query.Encode()
	}

	return uri
}

func (p Parameters) Set(key string, value string) {
	p[key] = value
}
