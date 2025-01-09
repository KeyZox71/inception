package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"strings"
	"bufio"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func removeSkipNetworking(filePath string) error {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a temporary slice to store updated lines
	var updatedLines []string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines that contain "skip-networking"
		if !strings.Contains(line, "skip-networking") {
			updatedLines = append(updatedLines, line)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Open the file for writing (overwrite mode)
	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	// Write the updated lines back to the file
	writer := bufio.NewWriter(file)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}
	return writer.Flush()
}

func createDBDir(dataDir string) {
	if dataDir == "/var/lib/mysql" {
		return
	}
	err := os.Mkdir(dataDir, 750)
	if err != nil {
		log.Fatal(err)
	}
}

func anyFileExists(folderPath string) (bool, error) {
	// Open the folder
	dir, err := os.Open(folderPath)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	// Read directory contents
	files, err := dir.Readdir(1) // Read at most 1 file
	if err != nil {
		return false, err
	}

	// If we read at least one file, it exists
	return len(files) > 0, nil
}

func checkOlderDB(dataDir string) bool {
	exist, err := anyFileExists(dataDir)
	if err != nil || exist == false {
		return false
	}
	return true
}

func waitForMariaDB(rootPass string) {
	for i := 0; i < 30; i++ {
		cmd := exec.Command("mariadb", "-uroot", "-p"+escapePassword(rootPass), "-e", "SELECT 1")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err == nil {
			fmt.Println("MariaDB is ready.")
			return
		} else {
			fmt.Printf("MariaDB init process in progress... (%d/%d): %v\n", i+1, 30, err)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("MariaDB init process failed.")
	os.Exit(1)
}

func escapeIdentifier(identifier string) string {
	// Replace backticks with double backticks to safely escape identifiers
	return fmt.Sprintf("`%s`", strings.ReplaceAll(identifier, "'", "\""))
}

func escapePassword(password string) string {
	// Escape single quotes in passwords
	return strings.ReplaceAll(password, "'", "\\'")
}

func configureMariaDB(rootPassword, user, password, database string) {
	cmd := exec.Command("mariadb", "-uroot", "-p"+rootPassword, "-e", fmt.Sprintf(`
		ALTER USER 'root'@'localhost' IDENTIFIED BY '%s';
		CREATE DATABASE IF NOT EXISTS %s;
		CREATE USER IF NOT EXISTS '%s'@'%%' IDENTIFIED BY '%s';
		GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';
		FLUSH PRIVILEGES;
	`, escapePassword(rootPassword), escapeIdentifier(database), escapeIdentifier(user), escapePassword(password), escapeIdentifier(database), escapeIdentifier(user)))
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
		filePath := "/etc/my.cnf.d/mariadb-server.cnf"

		oldDB := checkOlderDB(dataDir)
		if oldDB == false {
			createDBDir(dataDir)
			// Init DB
			cmd := exec.Command("mariadb-install-db", "--user=mysql", "--ldata="+dataDir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
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
			waitForMariaDB(rootPass)

			configureMariaDB(rootPass, user, pass, dbName)

			if err := removeSkipNetworking(filePath); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Successfully removed 'skip-networking' from the configuration file.")
			}

			cmd = exec.Command("mysqladmin", "-uroot", "-p"+rootPass, "shutdown")
			if err := cmd.Run(); err != nil {
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
