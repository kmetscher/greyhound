package main

import (
    "greyhound/core"
    "greyhound/routes"
)

func main() {
    router := core.Router()
    for _, route := range routes.Web() {
        router.Register(route)
    }
    router.Run(":8080")
}


