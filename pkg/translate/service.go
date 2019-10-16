package translate

import (
	"github.com/tentsk8s/tango/pkg/decoder"
	apps "k8s.io/api/apps/v1"
)

func ServiceToDeployment(svc *decoder.Service) *apps.Deployment {
	// https://godoc.org/k8s.io/api/apps/v1#Deployment
	return &apps.Deployment{
		// TODO: TypeMeta
		// https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta

		// TODO: ObjectMeta
		// https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta

		Spec: apps.DeploymentSpec{
			Replicas: intTo32(svc.Replicas),
			// TODO: LabelSelector
			// https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#LabelSelector

			// TODO: PodTemplateSpec
			// https://godoc.org/k8s.io/api/core/v1#PodTemplateSpec
			Template: containerToPodTemplateSpec(svc.Name, svc.Containers),

			// TODO: DeploymentStrategy
			// https://godoc.org/k8s.io/api/apps/v1#DeploymentStrategy

		},
	}
}
