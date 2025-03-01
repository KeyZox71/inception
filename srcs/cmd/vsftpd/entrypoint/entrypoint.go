package main

import (
	"os"

	"fmt"
	"log"
	"os/exec"

	_log "git.keyzox.me/42_adjoly/inception/internal/log"
)

func main() {
	args := os.Args

	if args[1] == "vsftpd" {
		_log.Log("note", "Entrypoint script for VSFTPD Server started")

		dir, err := os.ReadDir("/docker-entrypoint.d")
		if err != nil {
			log.Fatal(err)
		}
		_log.Log("note", "Running entrypoint scripts")
		for _, v := range dir {
			os.Chmod("/docker-entrypoint.d/"+v.Name(), 0755)
			cmd := exec.Command("/docker-entrypoint.d/" + v.Name())
			cmd.Env = os.Environ()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error running script(%s): %v\n", v.Name(), err)
				os.Exit(1)
			}
		}
	}
	cmd := exec.Command(args[1], args[2:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running NGINX: %v\n", err)
		os.Exit(1)
	}

}
