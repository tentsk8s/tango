package build

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	runner := &runner{}
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the app to get it ready to deploy",
		RunE:  runner.run,
	}
	flags := cmd.Flags()
	const dockerFileLocDescription = "The directory the Dockerfile is in"
	flags.StringVarP(&runner.dockerFileLoc, "dockerfile", "f", ".", dockerFileLocDescription)
	const dockerContextDescription = "The Docker context to send"
	flags.StringVarP(&runner.dockerContext, "context", "c", ".", dockerContextDescription)
	const buildTypeDescription = "The Docker builder type. 'local' for local Docker daemon, 'acr' for ACR builder"
	flags.StringVarP(&runner.builderStr, "builder", "b", "local", buildTypeDescription)
	const registryDescription = "The ACR registry (you must pass this if you pass build type as 'acr')"
	flags.StringVarP(&runner.registry, "registry", "r", "", registryDescription)

	return cmd
}
