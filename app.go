package onfig

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type App struct {
	pods        []pod
	deployments []deployment
}

func (a App) Pod(name, namespace string, deets *corev1.PodSpec) App {
	a.pods = append(a.pods, pod{
		name:      name,
		namespace: namespace,
		spec:      deets,
	})
	return a
}

func (a App) Deployment(name, namespace string, deets *appsv1.DeploymentSpec) App {
	a.deployments = append(a.deployments, deployment{
		name:      name,
		namespace: namespace,
		spec:      deets,
	})
}

func (a App) Run() error {
	for _, pod := range a.pods {
		// create pod
	}
	for _, deployment := range a.deployments {
		// create deployment
	}
	return nil
}
