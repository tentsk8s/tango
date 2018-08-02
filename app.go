package onfig

import (
	"github.com/tentsk8s/onfig/pod"
)

type App struct {
	pods []pod.Details
}

func (a App) Pod(deets *pod.Details) App {
	a.pods = append(a.pods)
	return a
}

func (a App) Run() error {
	return nil
}
