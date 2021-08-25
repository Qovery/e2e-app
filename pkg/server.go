package pkg

import (
	"fmt"
	"github.com/Qovery/e2e-app/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Route struct {
	Path string
	Method string
}

func serve(port string) {
	logrus.Infof("Starting server at port %s.", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logrus.Errorf("Can't create server: %s", err.Error())
		return
	}
}

func setRoutes(routes []Route) {
	for _, route := range routes {
		utils.CreateRoute(route.Path, route.Method)
		logrus.Infof("Route %s:%s created.", route.Method, route.Path)
	}
}

func CreateServer(routes []Route, port string) {
	setRoutes(routes)
	serve(port)
}