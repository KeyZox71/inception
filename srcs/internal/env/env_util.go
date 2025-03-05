package env

import (
	"os"
	"strings"

	"git.keyzox.me/42_adjoly/inception/internal/log"
)

func	removeNl(in string) string {

}

// Check if the "Value" exist in the env if not check if a "Value_FILE" exist and if still not just return default value
func	FileEnv(Value, Default string) string {
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

// Check if the "Value" exist in the env check if not just return default value
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
	} else if str[0] == '\'' && str[len(str) - 1] == '\'' {
		return strings.TrimPrefix(strings.TrimSuffix(str, "'"), "'")
	} else {
		return str
	}
}
