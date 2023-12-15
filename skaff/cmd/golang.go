package cmd

import (
	"github.com/STollenaar/AdventOfCode/skaff/golang"
	"github.com/spf13/cobra"
)

var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "Create scaffolding for a golang based day",
	RunE: func(cmd *cobra.Command, args []string) error {
		return datasource.Create(name, !clearComments, force)
	},
}

func init() {
	rootCmd.AddCommand(golangCmd)
	golangCmd.Flags().BoolVarP(&clearComments, "clear-comments", "c", false, "do not include instructional comments in source")
	golangCmd.Flags().StringVarP(&name, "name", "n", "", "name of the entity")
	golangCmd.Flags().BoolVarP(&force, "force", "f", false, "force creation, overwriting existing files")
}
