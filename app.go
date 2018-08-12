package main

import (
	"fmt"
	"goserver-usage/src/github.com/nekrasovdmytro/goserver"
	"net/http"
)

func main() {
	var routers []goserver.Router

	routerCollector := goserver.RouterCollector{
		Routers: routers,
	}

	routerCollector.AddRouter(goserver.Router{Path: "/hello", Func: func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}})

	server := goserver.Server{routerCollector}

	server.Run()
}
