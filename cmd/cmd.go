package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// baseCommand
type baseCommand struct {
	cmd *cobra.Command
}

// NewbaseCommand
func newbaseCommand(cmd *cobra.Command) *baseCommand {
	return &baseCommand{
		cmd: cmd,
	}
}

type cmder interface {
	getCommand() *cobra.Command
}

var rootCmd *baseCommand

func init() {
	rootCmd = newbaseCommand(
		&cobra.Command{
			Use:   "magazine",
			Short: "magazine",
			Long:  "magazine",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Usage()
			},
		},
	)
	rootCmd.cmd.PersistentFlags().StringP("author", "a", "", "author name")
}

func AddCommands(cmd cmder) {
	rootCmd.cmd.AddCommand(cmd.getCommand())
}

func Execute() {
	if err := rootCmd.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
