package routes

import (
    "greyhound/controllers"
)

// TODO: far more elegant route interface. This is confusing
func Web() []Route {
    webRoutes := []Route{
        {
            Method: Get,
            Path: "/",
            Controller: controllers.HelloWorld,
        },
        {
            Method: Get,
            Path: "/about",
            Controller: controllers.About,
        },
    }
    return webRoutes
}
