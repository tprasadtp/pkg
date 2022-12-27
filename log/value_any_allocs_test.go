package log_test

import (
	"net"
	"testing"
)

func TestAllocs(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")
	allocs := testing.AllocsPerRun(10, func() {
		ips := ip.String()
		_ = ips
	})
	if allocs != 0 {
		t.Errorf("%s => alloc mismatch actual=%f", t.Name(), allocs)
	}
}
