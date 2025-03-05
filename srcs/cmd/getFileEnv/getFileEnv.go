package main

import (
	"fmt"
	"os"

	"git.keyzox.me/42_adjoly/inception/internal/env"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		os.Exit(0)
	}

	env := env.FileEnv(args[1], "")
	if env == "" {
		os.Exit(1)
	}
	fmt.Print(env)
}
