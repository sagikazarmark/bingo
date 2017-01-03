package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GroupConstructor(t *testing.T) {
	group, _ := NewGroup("name", "description")

	assert.Equal(t, "name", group.Name)
	assert.Equal(t, "description", group.Description)
}

func Test_GroupConstructorErrorsIfNameIsEmpty(t *testing.T) {
	_, err := NewGroup("", "")

	assert.Error(t, err)
}

func Test_GroupAddEndpoint(t *testing.T) {
	group := &Group{}
	endpoint := &Endpoint{}

	group.AddEndpoint(endpoint)

	assert.Equal(t, endpoint, group.Endpoints[0])
}
