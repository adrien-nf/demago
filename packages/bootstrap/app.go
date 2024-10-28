package app

import (
	"demago/kernel"
	"demago/packages/routing"
	"fmt"
	"net/http"
)

type App struct{}

func (a *App) Start() {
	routerProvider := &kernel.RouterProvider{}
	router := routing.NewRouter()
	routerProvider.BootstrapRouter(router)
	
	fmt.Println("Server is running on :8080")
	
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
