package main

import (
	"bytes"
	"os"

	"fmt"
	"log"
	"os/exec"

	"git.keyzox.me/42_adjoly/inception/internal/cmd"
	"git.keyzox.me/42_adjoly/inception/internal/env"
	_log "git.keyzox.me/42_adjoly/inception/internal/log"
)

func	configFtp() {
	_log.Log("note", "Configuring VSFTPD...")
	ftpUser := env.FileEnv("FTP_USER", "ftp")
	ftpPass := env.FileEnv("FTP_PASS", "ftppass")
	cmd.ExecCmd([]string{"adduser", ftpUser, "--disabled-password"})

	var stdin bytes.Buffer
	stdin.WriteString(fmt.Sprintf("%s:%s", ftpUser, ftpPass))

	cmd := exec.Command("/usr/sbin/chpasswd")
	cmd.Stdin = &stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create("/etc/vsftpd.check")
	if err != nil {
		log.Fatal("could not create check file :O")
	}
	os.WriteFile("/etc/vsftpd/vsftpd.userlist", []byte(ftpUser), 0766)

	_log.Log("note", "VSFTPD configured ;D")
}

func main() {
	args := os.Args

	if args[1] == "vsftpd" {
		_log.Log("note", "Entrypoint script for VSFTPD Server started")

		_, err := os.ReadFile("/etc/vsftpd.check")
		if err != nil {
			configFtp()
		} else {
			_log.Log("note", "VSFTPD already configured, skipping...")
		}

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
		fmt.Printf("Error running VSFTPD: %v\n", err)
		os.Exit(1)
	}
}
