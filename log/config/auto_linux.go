//go:build linux

package config

import "github.com/tprasadtp/pkg/log"

func config(c AutomaticOptions) (log.Handler, error) {
	return &log.MockHandler{}, nil
}
