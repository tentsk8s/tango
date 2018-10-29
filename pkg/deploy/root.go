package deploy

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	runner := runner{}
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your app! Make sure to run 'tango build' before you do this",
		RunE:  runner.run,
	}
	return cmd
}
