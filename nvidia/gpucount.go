package nvidia

import (
	"os/exec"
	"strconv"
	"strings"
)

//GPUCount provides interface to get gpu count command and run it.
type GPUCount interface {
	command() *exec.Cmd
	run(cmd *exec.Cmd, env string) (int, error)
}

//Count implements one flavour of GPUCount interface.
type Count struct {
}

//NewCount returns instance of Count
func NewCount() Count {
	return Count{}
}

func (g Count) command() *exec.Cmd {
	cmd := "ls /dev | grep nvidia | grep -v nvidia-uvm | grep -v nvidiactl | wc -l"
	return exec.Command("bash", "-c", cmd)
}

func (g Count) run(cmd *exec.Cmd, env string) (int, error) {
	if env == "test" {
		return 4, nil
	}
	out, err := cmd.Output()
	ret := 0
	if err == nil {
		ret, _ = strconv.Atoi(strings.TrimSpace(string(out)))
		return ret, nil
	} else {
		return -1, err
	}
}
