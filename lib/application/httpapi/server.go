package httpapi

import (
	"fmt"
	"net/http"

	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
)

type HTTPServer struct {
	service input.IServiceFeature
	router  http.Handler
}

func NewHTTPServer(service input.IServiceFeature) *HTTPServer {
	res := new(HTTPServer)
	res.service = service
	res.router = res.newRouter()
	return res
}

func (me *HTTPServer) Start() {

	conf := config.GetConfiguration()

	logs.GetLogs().WriteMessage(logs.INFO, fmt.Sprintf("rubberyconf api started at port: %s", conf.Api.Port), nil)
	address := ":" + conf.Api.Port
	err := http.ListenAndServe(address, me.router)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("rubberyconf api error at port: %s", conf.Api.Port), err)
	}

}
