package middleware

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func isAsync(request http.Request) (async bool, e error) {
    if request.Header == nil {
        return false, errors.New("greyhound.core: Request has no headers")
    }
    if request.Header.Get("X-Greyhound-Async") == "true" {
        return true, nil
    }
    return false, nil
}


func Intercept() (handler gin.HandlerFunc) {

}

