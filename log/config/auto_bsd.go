//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package config

import "github.com/tprasadtp/pkg/log"

func config(c AutoConfigOptions) (log.Handler, error) {
	return &log.MockHandler{}, nil
}
