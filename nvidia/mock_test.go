package nvidia

import (
	"bufio"
	"os"
	"os/exec"
)

//MockLocal implements one flavour of Action interface.
type MockLocal struct {
}

//Start starts cmd and returns reader object that contains output of command
func (local MockLocal) start(cmd *exec.Cmd) *bufio.Reader {
	f, _ := os.Open("testing/gpuutil.csv")
	return bufio.NewReader(f)
}
