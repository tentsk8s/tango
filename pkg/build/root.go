package build

import (
	"github.com/spf13/cobra"
	"github.com/tentsk8s/tango/pkg/decoder"
	"github.com/tentsk8s/tango/pkg/log"
)

func Command() *cobra.Command {
	var dockerFileLoc string
	var dockerContext string
	var registry string
	var builderStr string
	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return log.Error(true, "Please don't pass any arguments to this :)")
			}
			builder, err := newBuilder(builderStr)
			if err != nil {
				return log.Err(true, err)
			}
			svc, err := decoder.Svc()
			if err != nil {
				return err
			}
			for _, container := range svc.Containers {
				log.Statusf(false, "Building image %s", container)
				if err := container.Validate(); err != nil {
					return log.Err(true, err)
				}

				dockerFile := container.GetDockerfile(dockerFileLoc)
				dockerRoot := container.GetDockerRoot(dockerContext)
				_, err := build(builder, buildParams{
					name:     container.Name,
					tag:      container.Tag,
					registry: registry,
					fileLoc:  dockerFile,
					context:  dockerRoot,
				})
				if err != nil {
					return err
				}

				// TODO: also push the image
			}
			return nil
		},
	}
	flags := cmd.Flags()
	const dockerFileLocDescription = "The directory the Dockerfile is in"
	flags.StringVarP(&dockerFileLoc, "dockerfile", "f", ".", dockerFileLocDescription)
	const dockerContextDescription = "The Docker context to send"
	flags.StringVarP(&dockerContext, "context", "c", ".", dockerContextDescription)
	const buildTypeDescription = "The Docker builder type. 'local' for local Docker daemon, 'acr' for ACR builder"
	flags.StringVarP(&builderStr, "builder", "b", "local", buildTypeDescription)
	const registryDescription = "The ACR registry (you must pass this if you pass build type as 'acr')"
	flags.StringVarP(&registry, "registry", "r", "", registryDescription)
	return cmd
}
