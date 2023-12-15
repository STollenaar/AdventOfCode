package cmd

import (
	"github.com/STollenaar/AdventOfCode/skaff/golang"
	"github.com/spf13/cobra"
)

var (
	name  string
	force bool
)

var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "Create scaffolding for a golang based day",
	RunE: func(cmd *cobra.Command, args []string) error {
		return golang.Create(name, force)
	},
}

func init() {
	rootCmd.AddCommand(golangCmd)
	golangCmd.Flags().StringVarP(&name, "name", "n", "", "name of the entity")
	golangCmd.Flags().BoolVarP(&force, "force", "f", false, "force creation, overwriting existing files")
}
