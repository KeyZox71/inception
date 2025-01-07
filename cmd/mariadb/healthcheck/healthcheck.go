package main

import (
	"database/sql"
	"fmt"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	"git.keyzox.me/42_adjoly/inception/internal/log"
)

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
	pass := env.FileEnv("MYSQL_PASSWORD", "default")
	user := env.FileEnv("MYSQL_USER", "mariadb")
	dbName := env.EnvCheck("MYSQL_DATABASE", "default")
	dbHost := "localhost"

	res := checkHealth(dbHost, user, pass, "3306", dbName)
	if res == true{
		_log.Log("note", "Mariadb is healthy")
	}
	_log.Log("error", "Health check failed")
}
