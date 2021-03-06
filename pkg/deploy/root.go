package deploy

import (
	"github.com/spf13/cobra"
	"github.com/tentsk8s/tango/pkg/decoder"
	"github.com/tentsk8s/tango/pkg/log"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return log.Error(true, "Don't pass any arguments to this command")
			}
			svc, err := decoder.Svc()
			if err != nil {
				return log.Err(true, err)
			}
			deployment := translate.ServiceToDeployment(svc)
			log.Statusf(true, "deployment %s created", deployment.Name)
			service := translate.ServiceToService(svc)

			return nil
		},
		// runner := &runner{}
		// cmd := &cobra.Command{
		// 	Use:   "deploy",
		// 	Short: "Deploy your app! Make sure to run 'tango build' before you do this",
		// 	RunE:  runner.run,
		// }
	}
	return cmd
}
