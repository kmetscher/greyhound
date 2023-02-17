package controllers

import (
    "github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) (component string, props interface{}) {
    helloProps := struct {
        Message string 
        Success string
    }{
        Message: "All aboard!",
        Success: "Your Greyhound app appears to be working.",
    }
    helloComponent := "HelloWorld"
    return helloComponent, helloProps
}
