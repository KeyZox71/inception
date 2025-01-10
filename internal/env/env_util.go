package env

import (
	"os"
	"strings"

	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func	FileEnv(Value string, Default string) string {
	val, is := os.LookupEnv(Value)

	if is {
		return val
	} else {
		val, is := os.LookupEnv(Value + "_FILE")
		if is {
			content, err := os.ReadFile(val)
			if err != nil {
				_log.Log("error", "Error reading file")
			}
			return string(content)
		}
	}
	return Default
}

func	EnvCheck(Value, Default string) string {
	val, is := os.LookupEnv(Value)

	if is {
		return val
	}
	return Default
}

func EscapeEnv(str string) string {
	if str[0] == '"' && str[len(str) - 1] == '"' {
		return strings.TrimPrefix(strings.TrimSuffix(str, "\""), "\"")
	} else if str[0] == '"' && str[len(str) - 1] == '"' {
		return strings.TrimPrefix(strings.TrimSuffix(str, "'"), "'")
	} else {
		return str
	}
}
