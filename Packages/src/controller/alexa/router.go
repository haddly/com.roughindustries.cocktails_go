// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/alexa/router.go: sets up all the routing for the alexa webapp
package alexa

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
	"net/http"
)

//Init to setup the http handlers
func AlexaRouterInit() {
	log.Println("Init in alexa/hello.go")
	router := mux.NewRouter()
	alexa.Init(Applications, router)

	n := negroni.Classic()
	n.UseHandler(router)
	http.Handle("/echo/helloworld", router)
}
