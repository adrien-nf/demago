package routing

import "net/http"

type RouteGroup struct {
    prefix string
    middlewares []string
    routes map[string]map[string]http.HandlerFunc
}

func NewRouteGroup() *RouteGroup {
    return &RouteGroup{routes: make(map[string]map[string]http.HandlerFunc), prefix: "/", middlewares: []string{}}
}

func (rg *RouteGroup) Get(path string, handler http.HandlerFunc) {
    if _, exists := rg.routes[path]; !exists {
        rg.routes[path] = make(map[string]http.HandlerFunc)
    }
    rg.routes[path]["GET"] = handler
}

func (rg *RouteGroup) Post(path string, handler http.HandlerFunc) {
    if _, exists := rg.routes[path]; !exists {
        rg.routes[path] = make(map[string]http.HandlerFunc)
    }
    rg.routes[path]["POST"] = handler
}

func (rg *RouteGroup) Put(path string, handler http.HandlerFunc) {
    if _, exists := rg.routes[path]; !exists {
        rg.routes[path] = make(map[string]http.HandlerFunc)
    }
    rg.routes[path]["PUT"] = handler
}

func (rg *RouteGroup) Patch(path string, handler http.HandlerFunc) {
    if _, exists := rg.routes[path]; !exists {
        rg.routes[path] = make(map[string]http.HandlerFunc)
    }
    rg.routes[path]["PATCH"] = handler
}

func (rg *RouteGroup) Delete(path string, handler http.HandlerFunc) {
    if _, exists := rg.routes[path]; !exists {
        rg.routes[path] = make(map[string]http.HandlerFunc)
    }
    rg.routes[path]["DELETE"] = handler
}

func (rg *RouteGroup) Prefix(prefix string) *RouteGroup {
    rg.prefix = prefix

    return rg
}

func (rg *RouteGroup) Middlewares(middlewares []string) *RouteGroup {
    rg.middlewares = middlewares

    return rg
}

func (rg *RouteGroup) Group(path string, handler func(r *RouteGroup)) *RouteGroup {
    r := NewRouteGroup()
    r.prefix = rg.prefix + path
    r.middlewares = rg.middlewares

    handler(r)

    for k, v := range r.routes {
        rg.routes[k] = v
    }

    return rg
}