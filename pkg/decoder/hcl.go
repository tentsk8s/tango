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

// GetDockerfile returns the path to the Dockerfile, not including the 'Dockerfile' part
func (c Container) GetDockerfile(def string) string {
	if c.Dockerfile == nil {
		return filepath.Join(def, "Dockerfile")
	}
	return filepath.Join(*c.Dockerfile, "Dockerfile")
}
func (c Container) GetDockerRoot(def string) string {
	if c.DockerRoot == nil {
		return def
	}
	return *c.DockerRoot
}

func (c Container) Validate() error {
	if !ioutil.IsFileExists(c.GetDockerfile(".")) {
		return fmt.Errorf("%s doesn't exist!", c.GetDockerfile("."))
	}
	if !ioutil.IsFileExists(c.GetDockerRoot(".")) {
		return fmt.Errorf("Docker context '%s' doesn't exist!", c.GetDockerRoot("."))
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
