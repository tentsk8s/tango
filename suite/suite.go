package suite

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tentsk8s/onfig"
	"github.com/tentsk8s/onfig/deployment"
	"github.com/tentsk8s/onfig/pod"
)

type Suite struct {
	suite.Suite
	app *onfig.App
	srv onfig.Server
}

func (s *Suite) BeforeTest() {
	a.srv = onfig.MockServer
	s.app.Server = a.srv
}

func (s Suite) Pods() []pod.Details {

}

func (s Suite) Deployments() []deployment.Details {

}

func Run(t *testing.T, s suite.TestingSuite) {
	suite.Run(t, s)
}