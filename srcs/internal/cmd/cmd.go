package cmd

import (
	"os"
	"os/exec"

	_log "git.keyzox.me/42_adjoly/inception/internal/log"
)

func ExecCmd(cmdStr []string) error {
	cmd := exec.Command(cmdStr[0], cmdStr...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
