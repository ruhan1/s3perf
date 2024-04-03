package main

import (
	"fmt"
	"os"

	cmd "github.com/ruhan1/s3perf/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "s3perf",
		Short: "tool to test performance for indy storage service",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(cmd.NewPrepareCmd())
	rootCmd.AddCommand(cmd.NewExecuteCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
