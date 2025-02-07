package main

import (
	"os"
	"fmt"
	"os/exec"

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
	fmt.Println("Maria	DB is healthy")
	return true
}

func main() {
	// Configuration
	user := env.EscapeEnv(env.FileEnv("MYSQL_USER", "mariadb"))
	password := env.EscapeEnv(env.FileEnv("MYSQL_PASSWORD", "default"))
	host := "127.0.0.1"
	port := "3306"

	if checkMariaDB(user, password, host, port) {
		os.Exit(0) // Success
	}

	fmt.Println("MariaDB health check failed")
	os.Exit(1) // Failure
}
