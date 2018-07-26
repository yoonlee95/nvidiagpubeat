package nvidia

import "testing"

func Test_GPUCount_Command(t *testing.T) {
	count := NewCount()
	cmd := count.command()

	if len(cmd.Args) != 3 {
		t.Errorf("Expected %d, Actual %d", 3, len(cmd.Args))
	}

	if cmd.Args[0] != "bash" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[0])
	}

	if cmd.Args[1] != "-c" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[1])
	}

	if cmd.Args[2] != "ls /dev | grep nvidia | grep -v nvidia-uvm | grep -v nvidiactl | wc -l" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[2])
	}
}

func Test_GPUCount_Run_TestEnv(t *testing.T) {
	count := NewCount()
	cmd := count.command()
	gpuCount, _ := count.run(cmd, "test")

	if gpuCount != 4 {
		t.Errorf("Expected %d, Actual %d", 4, gpuCount)
	}
}

func Test_GPUCount_Run_ProdEnv(t *testing.T) {
	count := NewCount()
	cmd := count.command()
	gpuCount, _ := count.run(cmd, "prod")

	if gpuCount != -1 {
		t.Errorf("Expected %d, Actual %d", -1, gpuCount)
	}
}
