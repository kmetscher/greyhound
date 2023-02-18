package controllers

import "github.com/gin-gonic/gin"

func About(c *gin.Context) (component string, props interface{}) {
    aboutProps := struct{
        Message string `json:"message"`
    }{
        Message: "Small miracle!",
    }
    aboutComponent := "Test"
    return aboutComponent, aboutProps
}
