package cmd

import (
	"os"
	"os/exec"
)

func ExecCmd(cmdStr, env []string) error {
	cmd := exec.Command(cmdStr[0], cmdStr[1:]...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
