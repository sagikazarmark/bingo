package main

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/julienschmidt/httprouter"
)

func Test_EndpointConstructor(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
	endpoint, _ := NewEndpoint("GET", "/endpoint", handler)

	assert.Equal(t, "GET", endpoint.Method)
	assert.Equal(t, "/endpoint", endpoint.Path)
}

func Test_EndpointConstructorErrorsIfMethodIsEmpty(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
	_, err := NewEndpoint("", "/endpoint", handler)

	assert.Error(t, err)
}

func Test_EndpointConstructorErrorsIfPathIsEmpty(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
	_, err := NewEndpoint("GET", "", handler)

	assert.Error(t, err)
}

func Test_EndpointConstructorErrorsIfHandlerIsEmpty(t *testing.T) {
	_, err := NewEndpoint("GET", "/endpoint", nil)

	assert.Error(t, err)
}

func Test_EndpointBuildUrl(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
	endpoint, _ := NewEndpoint("GET", "/endpoint/:param1/:param2", handler)

	endpoint.Parameters.Set("param1", "value1")
	endpoint.Parameters.Set("param2", "value2")

	assert.Equal(t, "/endpoint/value1/value2", endpoint.Url())
}

func Test_EndpointBuildUrlWithQuery(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
	endpoint, _ := NewEndpoint("GET", "/endpoint/:param1/:param2", handler)

	endpoint.Parameters.Set("param1", "value1")
	endpoint.Parameters.Set("param2", "value2")

	endpoint.Query.Set("param3", "value3")

	assert.Equal(t, "/endpoint/value1/value2?param3=value3", endpoint.Url())
}

func Test_SetParameters(t *testing.T) {
	parameters := make(Parameters)
	parameters.Set("key", "value")

	assert.Equal(t, "value", parameters["key"])
}
