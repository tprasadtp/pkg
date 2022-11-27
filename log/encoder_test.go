package log_test

import (
	"encoding/json"
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
)

func TestEventInterfaces(t *testing.T) {
	assert.Implements(t, (*json.Marshaler)(nil), new(log.Event), "log.Entry MarshalJSON")
}
