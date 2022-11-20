package version

import "testing"

func TestGetUserAgentSuffix(t *testing.T) {
	got := GetUserAgentSuffix()
	if got == "" {
		t.Error("got empty string for user agent")
	}
}

func TestGetUserAgent(t *testing.T) {
	got := GetUserAgent()
	if got == "" {
		t.Error("got empty string for user agent")
	}
}
