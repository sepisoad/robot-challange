package main

import (
	"log"

	"github.com/sepisoad/robot-challange/simulator/commands"
)

func main() {
	rootCmd := commands.Root()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
