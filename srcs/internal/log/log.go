package _log

import (
	"fmt"
	"os"
	"time"
)

func	Log(Type string, Log string) {
	t := time.Now()
	t.Format(time.RFC3339)
	if Type == "note" {
		fmt.Printf("%s-[%s] [Entrypoint]: %s\n", t, Type, Log)
	} else if Type == "warn" {
		fmt.Fprintf(os.Stderr ,"%s-[%s] [Entrypoint]: %s\n", t, Type, Log)
	} else if Type == "error" {
		fmt.Fprintf(os.Stderr ,"%s-[%s] [Entrypoint]: %s\n", t, Type, Log)
		os.Exit(1)
	}
}
