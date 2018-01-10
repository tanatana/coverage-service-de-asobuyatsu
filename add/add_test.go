package add_test

import (
	"testing"
	"github.com/tanatana/coverage-service-de-asobuyatsu/add"
)

func TestAdd(t *testing.T) {
	if add.Add(1, 1) != 2 {
		t.Error("error")
	}
	if add.Add(0, 1) != 1 {
		t.Error("error")
	}
	if add.Add(-1, 1) != 0 {
		t.Error("error")
	}
}
