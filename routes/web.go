package routes

import (
    "greyhound/controllers"
)

func Web() []Route {
    webRoutes := []Route{
        Route{
            Method: Get,
            Path: "/",
            Controller: controllers.HelloWorld,
        },
    }
    return webRoutes
}
