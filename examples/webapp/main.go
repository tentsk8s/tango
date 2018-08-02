package main

import (
	"github.com/tentsk8s/onfig"
)

func main() {
	a := app()
	a.Run()
}

func app() *onfig.App {
	a := onfig.App{}
	return a.WithDeployment(deployment.Details{
		Pod:      pod.Details{Image: "thing", Port: 8080},
		Replicas: 2,
		Name:"thing",
	}).WithService(service.Details{
		Port: 80,
		TargetPort: 8080,
	})
}
