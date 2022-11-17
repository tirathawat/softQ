package teardown_test

import (
	"testing"

	"github.com/tirathawat/softQ/pkg/teardown"
)

func TestTeardown(t *testing.T) {
	teardown := teardown.Setup(t)
	defer teardown()
}
