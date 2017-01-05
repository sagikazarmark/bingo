package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinConstructor(t *testing.T) {
	bin, _ := New("bingo", "Http integration tester", "Lorem ipsum")

	assert.Equal(t, "bingo", bin.Name)
	assert.Equal(t, "Http integration tester", bin.ShortDescription)
	assert.Equal(t, "Lorem ipsum", bin.Description)
}

func Test_BinConstructorErrorsIfNameIsEmpty(t *testing.T) {
	_, err := New("", "", "")

	assert.Error(t, err)
}

func Test_BinAddEndpoint(t *testing.T) {
	bin := &Bin{}
	endpoint := &Endpoint{}

	bin.AddEndpoint(endpoint)

	assert.Equal(t, endpoint, bin.Endpoints[0])
}

func Test_BinAddGroupt(t *testing.T) {
	bin := &Bin{}
	group := &Group{}

	bin.AddGroup(group)

	assert.Equal(t, group, bin.Groups[0])
}
