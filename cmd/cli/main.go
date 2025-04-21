package main

import (
	"fmt"
	"os"

	"github.com/gblcarvalho/go-expert-stress-test/internal/commands"
)

func main() {
	cmd := commands.NewStressTestCMD()
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(1)
	}
}
