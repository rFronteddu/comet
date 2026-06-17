package version

import (
	"strings"
	"testing"
)

func TestBannerIncludesNameAndVersion(t *testing.T) {
	banner := Banner()

	if !strings.Contains(banner, "COMET") {
		t.Fatal("banner does not include COMET")
	}

	if !strings.Contains(banner, Version) {
		t.Fatal("banner does not include version")
	}
}
