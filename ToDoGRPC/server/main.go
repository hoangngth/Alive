package main

import (
	"fmt"
	"os"

	"Alive/ToDoGRPC/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
