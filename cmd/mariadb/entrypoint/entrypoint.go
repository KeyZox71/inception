package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func createDBDir(dataDir string) {
	if dataDir == "/var/lib/mysql" {
		return
	}
	err := os.Mkdir(dataDir, 750)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(dataDir+"/.already", 750)
	if err != nil {
		log.Fatal(err)
	}
}

func checkOlderDB(dataDir string) bool {
	if _, err := os.Stat(dataDir+"/.already"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func waitForMariaDB() {
	for i := 0; i < 30; i++ {
		cmd := exec.Command("mysql", "-uroot", "-e", "SELECT 1")
		if err := cmd.Run(); err == nil {
			return
		}
		fmt.Println("MariaDB init process in progress...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("MariaDB init process failed.")
	os.Exit(1)
}

func configureMariaDB(rootPassword, user, password, database string) {
	cmd := exec.Command("mysql", "-uroot", "-e", fmt.Sprintf(`
		ALTER USER 'root'@'localhost' IDENTIFIED BY '%s';
		CREATE DATABASE IF NOT EXISTS %s;
		CREATE USER IF NOT EXISTS '%s'@'%%' IDENTIFIED BY '%s';
		GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';
		FLUSH PRIVILEGES;
	`, rootPassword, database, user, password, database, user))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error configuring MariaDB: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args


	if args[1] == "mariadbd" || args[1] == "mysqld" {
		_log.Log("note", "Entrypoint script for MariaDB Server started")

		rootPass := env.FileEnv("MYSQL_ROOT_PASSWORD", "default")
		pass := env.FileEnv("MYSQL_PASSWORD", "default")
		user := env.FileEnv("MYSQL_USER", "mariadb")
		dbName := env.EnvCheck("MYSQL_DATABASE", "default")
		dataDir := env.EnvCheck("DATADIR", "/var/lib/mysql")

		oldDB := checkOlderDB(dataDir)
		if oldDB == false {
			createDBDir(dataDir)
			// Init DB
			cmd := exec.Command("mariadb-install-db", "--user=mysql", "--ldata="+dataDir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			createDBDir(dataDir)
			if err := cmd.Run(); err != nil {
				_log.Log("error", "Error initializing MariaDB")
			}

			// Starting temp mariadb server for config
			cmd = exec.Command("mariadbd-safe", "--skip-networking")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				fmt.Printf("Error starting MariaDB: %v\n", err)
				os.Exit(1)
			}
			// Wait for mariadb to start
			waitForMariaDB()

			configureMariaDB(rootPass, user, pass, dbName)

			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Error stopping MariaDB: %v\n", err)
			}
		}
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running MariaDB: %v\n", err)
		os.Exit(1)
	}
}
