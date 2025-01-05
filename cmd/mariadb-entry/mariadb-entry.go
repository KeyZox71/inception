package main

import (
	"fmt"
	"os"
)

func	main() {
	args := os.Args	

	for _, v := range args {
		fmt.Println(v)
	}
}
