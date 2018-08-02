package onfig

import (
	"k8s.io/api/core/v1"
)

type pod struct {
	name      string
	namespace string
	spec      *v1.PodSpec
}
