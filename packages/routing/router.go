package routing

import "net/http"

type Router struct {
    routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
    return &Router{routes: make(map[string]map[string]http.HandlerFunc)}
}

func (r *Router) Handle(routeGroup *RouteGroup) {
    for path, methodHandlers := range routeGroup.routes {
        if _, exists := r.routes[path]; !exists {
            r.routes[path] = make(map[string]http.HandlerFunc)
        }

        for method, handler := range methodHandlers {
            r.routes[path][method] = handler
        }
    }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    if handlers, exists := r.routes[req.URL.Path]; exists {
        if handler, methodExists := handlers[req.Method]; methodExists {
            handler(w, req)
            return
        }
    }
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
