package kernel

import (
	routing "demago/packages/routing"
	api "demago/routes/api"
	web "demago/routes/web"
)

type RouterProvider struct{}

func (r *RouterProvider) BootstrapRouter(router *routing.Router) {
    router.Handle(web.SetupRoutes(routing.NewRouteGroup()))

    router.Handle(api.SetupRoutes(routing.NewRouteGroup().Prefix("/api").Middlewares([]string{"auth", "api"})))
}
