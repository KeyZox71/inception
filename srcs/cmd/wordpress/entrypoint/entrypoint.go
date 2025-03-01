package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func makeFpmConf() {
	val, is := os.LookupEnv("PHP_NOT_CONFIGURE")
	_, err := os.ReadFile("/www-docker.conf")

	if (is == true && val == "true") {
		_log.Log("note", "PHP-FPM - Not configuring (PHP-CONFIGURE set to true)") 
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		_log.Log("note", "PHP-FPM - already configured, skipping") 
		return
	}
	_log.Log("note", "PHP-FPM - Configuring...")
	v, is := os.LookupEnv("PHP_PORT")
	content, err := os.ReadFile("/www-docker.conf")
	if err != nil {
		log.Fatal(err)
	}
	if !is {
		v = "9000"
	}
	res := bytes.Replace([]byte(content), []byte("$PHP_PORT"), []byte(v), -1)

	if err := os.WriteFile("/etc/php84/php-fpm.d/www.conf", res, 0660); err != nil {
		log.Fatal(err)
	}
	os.Remove("/www-docker.conf")
}

func main() {
	args := os.Args

	if args[1] == "php-fpm84" {
		_log.Log("note", "Entrypoint script for Wordpress-PHP-FPM Server started")

		makeFpmConf()

		dir, err := os.ReadDir("/docker-entrypoint.d")
		if err != nil {
			log.Fatal(err)
		}
		_log.Log("note", "Running entrypoint scripts...")
		for _, v := range dir {
			os.Chmod("/docker-entrypoint.d/" + v.Name(), 0755)
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

	_log.Log("note", "Starting container")
	cmd := exec.Command(args[1], args[2:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running MariaDB: %v\n", err)
		os.Exit(1)
	}
}
