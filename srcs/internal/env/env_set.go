package env

import (
	"log"
	"os"
	"regexp"
)

func IsEnvSet(what string) bool {
	env := os.Environ()

	reg, err := regexp.Compile("^" + what)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range env {
		if reg.MatchString(v) {
			return true
		}
	}
	return false
}
