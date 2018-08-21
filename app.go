package main

import (
	"goserver-usage/src/github.com/nekrasovdmytro/goserver"
	"net/http"
	"goserver-usage/src/github.com/nekrasovdmytro/goconfig"
)

func main() {
	var routers []goserver.Router

	routerCollector := goserver.RouterCollector{
		Routers: routers,
	}

	routerCollector.AddRouter(goserver.Router{Method: http.MethodGet, Path: "/", Func: func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]string)

		data["element-1"] = "value-1"

		message := goserver.JsonMessage{
			Code: 0,
			Message: "",
			Data: data,
		}

		goserver.HandelJsonResponse(w, http.StatusOK, message)
	}})

	configurator := goconfig.Configurator{}
	configurator.SetLoader(goconfig.JsonConfigLoader{})
	configurator.Load("config/goserver.json")

	configData := configurator.GetData()
	configDataMap := configData.(map[string]string)

	config := goserver.ServerConfig{
		Host: configDataMap["Host"],
		Port: configDataMap["Port"],
	}

	server := goserver.Server{RouterCollector : routerCollector, ServerConfig: config}

	error := server.Run()

	if error != nil {
		panic(error)
	}
}
