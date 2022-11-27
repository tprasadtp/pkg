package log_test

import (
	"encoding/json"
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/log"
)

func TestEntry_Implements(t *testing.T) {
	assert.Implements(t, (*json.Marshaler)(nil), new(log.Entry), "log.Entry MarshalJSON")
}
