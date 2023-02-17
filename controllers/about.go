package controllers

import "github.com/gin-gonic/gin"

func About(c *gin.Context) (component string, props interface{}) {
    aboutComponent := "About"
    aboutProps := struct{
        msg string 
        link string
    }{
        msg: "Greyhound is a project by Kyle Metscher",
        link: "https://kylemetscher.com",
    }
    return aboutComponent, aboutProps
}
