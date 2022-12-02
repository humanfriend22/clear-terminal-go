package terminal

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	ClearCustom(os.Stdout)
}

func ClearCustom(stdout *os.File) {
	command := "clear"

	if runtime.GOOS == "win32" {
		command = "cls"
	}

	cmd := exec.Command(command)
	cmd.Stdout = stdout
	cmd.Run()
}
