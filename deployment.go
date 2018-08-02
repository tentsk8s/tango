package onfig

import (
	"k8s.io/api/apps/v1"
)

type deployment struct {
	name      string
	namespace string
	spec      *v1.DeploymentSpec
}
