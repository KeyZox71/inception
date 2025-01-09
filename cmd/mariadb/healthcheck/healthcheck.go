package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func escapeIdentifier(identifier string) string {
	// Replace backticks with double backticks to safely escape identifiers
	return fmt.Sprintf("`%s`", strings.ReplaceAll(identifier, "'", "\""))
}

func escapePassword(password string) string {
	// Escape single quotes in passwords
	return strings.ReplaceAll(password, "'", "\\'")
}

func checkHealth(host, user, pass, port, dbName string) bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)

	// Attempt to open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		_log.Log("warning", fmt.Sprintf("Failed to open database connection: %v", err))
		return false
	}
	defer db.Close()

	// Attempt to ping the database
	if err := db.Ping(); err != nil {
		_log.Log("warning", fmt.Sprintf("Health check failed: %v", err))
		return false
	}

	_log.Log("note", "Health check passed successfully")
	return true
}

func main() {
	// Load environment variables
	pass := escapePassword(env.FileEnv("MYSQL_PASSWORD", "default"))
	user := escapeIdentifier(env.FileEnv("MYSQL_USER", "mariadb"))
	dbName := escapeIdentifier(env.EnvCheck("MYSQL_DATABASE", "default"))
	dbHost := "127.0.0.1"

	// Perform the health check
	res := checkHealth(dbHost, user, pass, "3306", dbName)
	if res {
		_log.Log("note", "MariaDB is healthy")
		os.Exit(0)
	}

	_log.Log("warning", "Health check failed")
	os.Exit(1)
}
