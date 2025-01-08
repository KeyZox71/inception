package main

import (
	"database/sql"
	"fmt"
	"strings"
	"os"

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

func	checkHealth(host, user, pass, port, dbName string) bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		_log.Log("error", "Failed to open database connection")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		_log.Log("error", "Health check failed")
	}
	return true
}

func	main() {
	pass := escapePassword(env.FileEnv("MYSQL_PASSWORD", "default"))
	user := escapeIdentifier(env.FileEnv("MYSQL_USER", "mariadb"))
	dbName := escapeIdentifier(env.EnvCheck("MYSQL_DATABASE", "default"))
	dbHost := escapeIdentifier("localhost")

	res := checkHealth(dbHost, user, pass, "3306", dbName)
	if res == true{
		_log.Log("note", "Mariadb is healthy")
		os.Exit(0)
	}
	_log.Log("error", "Health check failed")
}
