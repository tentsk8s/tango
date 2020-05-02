package server

import (
	"github.com/spf13/cobra"
)

type runner struct{}

func (r *runner) run(cmd *cobra.Command, args []string) error { return nil }
