package main

import (
	"fmt"

	"git.keyzox.me/42_adjoly/inception/internal/env"
	_log "git.keyzox.me/42_adjoly/inception/internal/log"
)

func main(){
	pass := env.FileEnv("BORG_PASSPHRASE", "")
	if pass == "" {
		_log.Log("error", "Could not found passphrase")
	}
	fmt.Print(pass)
}
