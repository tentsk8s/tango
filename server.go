package onfig

type Server interface {
	Send(a *App)
}

type MockServer struct {
	// stuff
}

// TODO: implement Send()
