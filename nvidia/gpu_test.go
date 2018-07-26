package nvidia

import (
	"fmt"
	"os"
	"testing"
)

func Test_Command_TestEnv(t *testing.T) {
	util := NewUtilization()
	cmd := util.command("test", "myquery")

	if len(cmd.Args) != 1 {
		t.Errorf("Expected %d, Actual %d", 1, len(cmd.Args))
	}

	if cmd.Args[0] != "localnvidiasmi" {
		t.Errorf("Expected %s, Actual %s", "localnvidiasmi", cmd.Args[0])
	}
}

func Test_Command_ProdEnv(t *testing.T) {
	util := NewUtilization()
	cmd := util.command("prod", "myquery")

	if len(cmd.Args) != 3 {
		t.Errorf("Expected %d, Actual %d", 3, len(cmd.Args))
	}

	if cmd.Args[0] != "nvidia-smi" {
		t.Errorf("Expected %s, Actual %s", "nvidia-smi", cmd.Args[0])
	}

	if cmd.Args[1] != "--query-gpu=myquery" {
		t.Errorf("Expected %s, Actual %s", "--query-gpu=myquery", cmd.Args[0])
	}

	if cmd.Args[2] != "--format=csv" {
		t.Errorf("Expected %s, Actual %s", "--format=csv", cmd.Args[0])
	}
}

func Test_Run_TestEnv(t *testing.T) {
	util := NewUtilization()
	query := "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate"
	cmd := util.command("test", query)
	os.Setenv("PATH", ".")
	output, _ := util.run(cmd, 4, query, NewLocal())
	if output == nil {
		//TODO fix unit test case
		//t.Errorf("output cannot be nil")
	}
}

func Test_Run_ProdEnv(t *testing.T) {
	util := NewUtilization()
	query := "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate"
	cmd := util.command("prod", query)
	fmt.Println(cmd.Path)
	output, _ := util.run(cmd, 4, query, MockLocal{})

	for _, o := range output {
		if o == nil {
			t.Errorf("output cannot be nil.")
		}
	}

}
