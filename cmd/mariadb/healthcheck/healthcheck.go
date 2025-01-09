package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"git.keyzox.me/42_adjoly/inception/internal/env"
)

func checkMariaDB(user, password, host, port string) bool {
	// Create the command to run mariadb client
	cmd := exec.Command("mariadb", "-u"+user, "-p"+password, "-h"+host, "-P"+port, "-e", "SELECT 1;")
	
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Health check failed: %v\n", err)
		return false
	}
	fmt.Println("MariaDB is healthy")
	return true
}

func escapePassword(password string) string {
	// Escape single quotes in passwords
	new := strings.ReplaceAll(password, "\"", "")
	return strings.ReplaceAll(new, "'", "\\'")
}

func main() {
	// Configuration
	user := escapePassword(env.FileEnv("MYSQL_USER", "mariadb"))
	password := escapePassword(env.FileEnv("MYSQL_PASSWORD", "default"))
	host := "127.0.0.1"
	port := "3306"

	// Retry health check for MariaDB
	for i := 0; i < 10; i++ {
		if checkMariaDB(user, password, host, port) {
			os.Exit(0) // Success
		}
		fmt.Println("Waiting for MariaDB to become ready...")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("MariaDB health check failed")
	os.Exit(1) // Failure
}
