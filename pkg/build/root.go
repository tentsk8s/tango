package build

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
	"github.com/spf13/cobra"
	"github.com/tentsk8s/tango/pkg/decoder"
	"github.com/tentsk8s/tango/pkg/log"
)

func buildImage(imageName, imageTag, dockerFileLoc, buildContextLoc string) (bool, error) {
	img := fmt.Sprintf("%s:%s", imageName, imageTag)
	args := []string{"build", "-t", img}
	if dockerFileLoc != "." {
		args = append(args, "-f", dockerFileLoc)
	}
	args = append(args, buildContextLoc)
	return sh.Exec(map[string]string{}, os.Stdout, os.Stderr, "docker", args...)
	// TODO: also push the container
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return log.Error(true, "Please don't pass any arguments to this :)")
			}
			svc, err := decoder.Svc()
			if err != nil {
				return err
			}
			for _, container := range svc.Containers {
				log.Statusf(false, "Building %s...", container)
				if err := container.Validate(); err != nil {
					return log.Err(true, err)
				}
				_, err := buildImage(container.Name, container.Tag, container.GetDockerfile(), container.GetDockerRoot())
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
