package routing

import (
	"net/http"
)

type Route struct {
    routes map[string]map[string]http.HandlerFunc
}

func NewRoute() *Route {
    return &Route{
        routes: make(map[string]map[string]http.HandlerFunc),
    }
}
