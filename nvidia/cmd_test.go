package nvidia

import (
	"os/exec"
	"testing"
)

func Test_start(t *testing.T) {
	cmd := exec.Command("ls")

	local := NewLocal()
	reader := local.start(cmd)

	if reader == nil {
		t.Errorf("Reader cannot be nil")
	}
}
