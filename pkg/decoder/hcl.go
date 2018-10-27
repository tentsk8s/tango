package decoder

import (
	"fmt"
	"path/filepath"

	"github.com/appscode/go/ioutil"
)

type Service struct {
	Name       string
	Replicas   int
	Containers []Container
}

type Container struct {
	Name       string
	Tag        string
	Port       int
	Env        []string
	Dockerfile *string
	DockerRoot *string
}

func (c Container) GetDockerfile() string {
	if c.Dockerfile == nil {
		return "."
	}
	return *c.Dockerfile
}
func (c Container) GetDockerRoot() string {
	if c.DockerRoot == nil {
		return "."
	}
	return *c.DockerRoot
}

func (c Container) Validate() error {
	dockerFileFull := filepath.Join(c.GetDockerfile(), "Dockerfile")
	if !ioutil.IsFileExists(dockerFileFull) {
		return fmt.Errorf("%s doesn't exist!", dockerFileFull)
	}
	return nil
}

func (c Container) CompleteImageName() string {
	return c.Name + ":" + c.Tag
}

func (c Container) String() string {
	return c.CompleteImageName()
}

func Svc() (*Service, error) {
	return &Service{
		Name: "sample",
		Containers: []Container{
			{Name: "tentsk8s/sample", Tag: "latest"},
		},
	}, nil
}
