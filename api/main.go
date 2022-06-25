package main

import (
	"log"

	"github.com/sepisoad/robot-challange/api/commands"
)

func main() {
	rootCmd := commands.Root()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
