package clicolors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummyUnix(t *testing.T) {
	assert.Nil(t, nil, EnableVTProcessing())
}
