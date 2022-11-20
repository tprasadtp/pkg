package log_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"github.com/tprasadtp/pkg/log"
)

func TestFromContext(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	logger := log.FromContext(ctx)
	is.Equal(log.Log, logger)

	logs := log.WithField("foo", "bar")
	ctx = log.NewContext(ctx, logs)

	logger = log.FromContext(ctx)
	is.Equal(logs, logger)
}
