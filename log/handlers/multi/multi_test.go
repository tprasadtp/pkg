package multi_test

import (
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
	"github.com/tprasadtp/pkg/log/handlers/discard"
	"github.com/tprasadtp/pkg/log/handlers/multi"
)

func TestInterface(t *testing.T) {
	assert.Implements(t,
		(*log.Handler)(nil),
		multi.New(discard.New(log.DebugLevel), discard.New(log.InfoLevel)),
		"multi.Handler => log.Handler")
}
