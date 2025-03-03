package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"bufio"

	"git.keyzox.me/42_adjoly/inception/internal/cmd"
	"git.keyzox.me/42_adjoly/inception/internal/env"
	_log "git.keyzox.me/42_adjoly/inception/internal/log"
	"git.keyzox.me/42_adjoly/inception/internal/pass"
)

func overrideCronFile(filePath string, jobs []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, job := range jobs {
		_, err := writer.WriteString(job + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func isBorgInit(repo string) (bool, error) {
	cmd := exec.Command("borg", "info", repo)

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 2 {
				return false, nil
			}
		}
		return false, err
	}
	return true, err
}

func main() {
	args := os.Args

	src := env.EnvCheck("BORG_SRC", "/source")
	if _, err := os.ReadDir(src); err != nil {
		_log.Log("error", src+" does not exist can't perform backup")
	}

	repo := env.EnvCheck("BORG_REPO", "/backup")
	if _, err := os.ReadDir(src); err != nil {
		_log.Log("error", repo+" does not exist can't perform backup")
	}
	is, err := isBorgInit(repo)
	if err != nil {
		log.Fatal(err)
	} else if is == true {
		_log.Log("note", "Repo already initialize, skipping...")
	} else {
		_log.Log("note", "Initializing repo...")

		passphrase := env.FileEnv("BORG_PASSPHRASE", "")
		if passphrase == "" {
			_log.Log("error", "No passphrase specified, exiting...")
		}

		err = cmd.ExecCmd([]string{"borg", "init", "--encryption=" + passphrase, repo})
		if err != nil {
			log.Fatal(err)
		}
	}

	interval := env.EnvCheck("CRON_INTERVAL", "0 0 * * *")
	cronFilePath := "/etc/crontabs/root"
	newJobs := []string{
		"# Borg Backup Cron Job",
		interval + " root run-parts /docker-backup.d >> /var/log/cron.log 2>&1",
	}
	err = overrideCronFile(cronFilePath, newJobs)
	if err != nil {
		fmt.Printf("Error overriding cron file: %v\n", err)
	} else {
		fmt.Println("Cron file overridden successfully.")
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

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running BORGBACKUP: %v\n", err)
		os.Exit(1)
	}
}
