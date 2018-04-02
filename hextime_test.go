package hextime_test

import (
	"testing"

	"github.com/mcandre/go-hextime"
)

func TestVersion(t *testing.T) {
	if hextime.Version == "" {
		t.Errorf("Expected %s to be non-blank", hextime.Version)
	}
}
