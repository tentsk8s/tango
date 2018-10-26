package deploy

import (
	"log"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				log.Fatal("don't pass any arguments to this")
			}
			return nil
		},
	}
}
