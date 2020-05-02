package build

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

type dockerBuilder string

const (
	localDaemonBuilder dockerBuilder = "local"
	acrBuilder         dockerBuilder = "acr"
)

func newBuilder(s string) (dockerBuilder, error) {
	switch s {
	case string(localDaemonBuilder):
		return localDaemonBuilder, nil
	case string(acrBuilder):
		return acrBuilder, nil
	default:
		return "", fmt.Errorf("'%s' is an unsupported builder", s)
	}
}

type buildParams struct {
	name     string
	tag      string
	registry string
	fileLoc  string
	context  string
}

func (b buildParams) fullImage() string {
	return fmt.Sprintf("%s:%s", b.name, b.tag)
}

func build(builder dockerBuilder, params buildParams) (bool, error) {
	if builder == localDaemonBuilder {
		return dockerBuild(params)
	}
	return acrBuild(params)
}

func dockerBuild(params buildParams) (bool, error) {
	img := params.fullImage()
	args := []string{"build", "-t", img}
	if params.fileLoc != "." {
		args = append(args, "-f", params.fileLoc)
	}
	args = append(args, params.context)
	return sh.Exec(map[string]string{}, os.Stdout, os.Stderr, "docker", args...)

}

func acrBuild(params buildParams) (bool, error) {
	if params.registry == "" {
		return false, fmt.Errorf("registry must be set for ACR builds")
	}
	img := params.fullImage()
	args := []string{"acr", "build", "-t", img, "-r", params.registry}
	if params.fileLoc != "." {
		args = append(args, "-f", params.fileLoc)
	}
	args = append(args, params.context)
	return sh.Exec(map[string]string{}, os.Stdout, os.Stderr, "az", args...)
}
