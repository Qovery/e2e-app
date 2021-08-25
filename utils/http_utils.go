package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type CallBack func(http.ResponseWriter, *http.Request)

type ResponseJson struct {
	Message string `json:"message"`
}

func createCallback(path string, method string) CallBack {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/%s", path) {
			http.Error(w, "404 not found.", http.StatusNotFound)
			logrus.Errorf("Can't find path %s", r.URL.Path)
			return
		}

		if r.Method != strings.ToUpper(method) {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(ResponseJson{fmt.Sprintf("OK from %s", path)})
	}
}

func CreateRoute(path string, method string) {
	callBack := createCallback(path, method)
	http.HandleFunc(fmt.Sprintf("/%s", path), callBack)
}
