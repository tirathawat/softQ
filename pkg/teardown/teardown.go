package teardown

import "testing"

func Setup(t *testing.T) func() {
	t.Log("Setup")

	return func() {
		t.Log("Teardown")
	}
}
