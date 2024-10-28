package api

import (
	"demago/packages/routing"
	"fmt"
	"net/http"
)

func SetupRoutes(r *routing.RouteGroup) *routing.RouteGroup {
    r.Get("/pizzas", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "GET request for pizzas")
    })

    r.Post("/pizzas", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "POST request for pizzas")
    })

    r.Patch("/pizzas", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "UPDATE request for pizzas")
    })

    return r
}