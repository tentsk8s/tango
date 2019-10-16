package server

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	runner := &runner{}
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run the tango server",
		RunE:  runner.run,
	}

	return cmd
}
