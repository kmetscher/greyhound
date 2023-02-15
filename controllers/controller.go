package controllers

import (
    "github.com/gin-gonic/gin"
)

type Controller func(c *gin.Context) (component string, props interface{})
