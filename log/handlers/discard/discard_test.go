package discard_test

import (
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*log.Handler)(nil), discard.New(log.DEBUG), "discard.Handler => log.Handler")
}
