package controllers

import (
    "github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) (component string, props interface{}) {
    helloProps := struct {
        msg string 
    }{
        msg: "All aboard!",
    }
    helloComponent := "HelloWorld"
    return helloComponent, helloProps
}
