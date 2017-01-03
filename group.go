package bingo

import "errors"

// Holds a list of endpoints
type Group struct {
	Name        string
	Description string
	Endpoints   []*Endpoint
}

func NewGroup(name string, description string) (*Group, error) {
	if name == "" {
		return nil, errors.New("Name must not be empty")
	}

	return &Group{
		Name:        name,
		Description: description,
	}, nil
}

// Idiomatic way of adding an Endpoint without initializing the underlying slice
func (g *Group) AddEndpoint(endpoint *Endpoint) {
	g.Endpoints = append(g.Endpoints, endpoint)
}
