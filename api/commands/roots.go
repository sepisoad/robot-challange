package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCommand returns root CLI application command interface
func Root() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "robots-api",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// Register all commands
	cmd.AddCommand(
		startCommand(),
	)

	return cmd
}

func printHeader() {
	fmt.Println("Robots API")
}
