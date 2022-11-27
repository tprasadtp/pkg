package noop_test

import (
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/noop"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*log.Handler)(nil), noop.New(log.DEBUG), "noop.Handler => log.Handler")
}
