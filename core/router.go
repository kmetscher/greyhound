package core

import (
    "github.com/gin-gonic/gin"
    "greyhound/middleware"
)

type Greyhound struct {
    engine gin.Engine
}

func New() *Greyhound {
    engine := gin.Default()
    engine.Use(middleware.Intercept())
    g := &Greyhound{
        *engine,
    }
    return g
}
