package build

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tentsk8s/tango/pkg/decoder"
	"github.com/tentsk8s/tango/pkg/git"
	"github.com/tentsk8s/tango/pkg/log"
)

type runner struct {
	dockerFileLoc string
	dockerContext string
	registry      string
	builderStr    string
}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("Please don't pass any arguments to this command :)")
	}

	builder, err := newBuilder(r.builderStr)
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
		container.Tag, err = git.SHA()
		if err != nil {
			return log.Errorf(true, "couldn't get git SHA for container tag")
		}

		dockerFile := container.GetDockerfile(r.dockerFileLoc)
		dockerRoot := container.GetDockerRoot(r.dockerContext)
		_, err := build(builder, buildParams{
			name:     container.Name,
			tag:      container.Tag,
			registry: r.registry,
			fileLoc:  dockerFile,
			context:  dockerRoot,
		})
		if err != nil {
			return err
		}

		// TODO: also push the image
		// TODO: write the container back to list
	}

	// TODO: write the container list back to the file
	return nil
}
