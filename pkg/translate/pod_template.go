package translate

import (
	"fmt"

	"github.com/tentsk8s/tango/pkg/decoder"
	core "k8s.io/api/core/v1"
)

// https://godoc.org/k8s.io/api/core/v1#PodTemplateSpec
func containerToPodTemplateSpec(svcName string, decoderContainers []decoder.Container) core.PodTemplateSpec {
	// https://godoc.org/k8s.io/api/core/v1#Container
	containers := make([]core.Container, len(decoderContainers))
	for i, decoderContainer := range decoderContainers {
		containers[i] = core.Container{
			Name:  fmt.Sprintf("%s-%d", svcName, i),
			Image: decoderContainer.CompleteImageName(),
			// https://godoc.org/k8s.io/api/core/v1#ContainerPort
			Ports: []core.ContainerPort{
				{ContainerPort: int32(decoderContainer.Port)},
			},
			// TODO: Env
			// https://godoc.org/k8s.io/api/core/v1#EnvVar
			ImagePullPolicy: core.PullAlways,
		}
	}
	return core.PodTemplateSpec{
		// TODO: ObjectMeta
		// https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta

		// https://godoc.org/k8s.io/api/core/v1#PodSpec
		Spec: core.PodSpec{
			Containers: containers,
		},
	}
}
