package nvidia

import (
	"bufio"
	"os/exec"
)

//Action provides interface to start execution of a command
type Action interface {
	start(cmd *exec.Cmd) *bufio.Reader
}

//Local implements one flavour of Action interface.
type Local struct {
}

//NewLocal returns instance of Local
func NewLocal() Local {
	return Local{}
}

//Start starts cmd and returns reader object that contains output of command
func (local Local) start(cmd *exec.Cmd) *bufio.Reader {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	return bufio.NewReader(stdout)
}
