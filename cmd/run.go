package cmd

import "github.com/Qovery/e2e-app/pkg"

func Execute(port string) {
	routes := []pkg.Route{
		{"test", "get"},
	}
	pkg.CreateServer(routes, port)
}
