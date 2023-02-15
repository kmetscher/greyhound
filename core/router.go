package core

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"greyhound/bootstrap"
	"greyhound/controllers"
	"greyhound/routes"
)

type Greyhound struct {
    engine gin.Engine
}

func New() *Greyhound {
    engine := gin.Default()
    engine.SetHTMLTemplate(bootstrap.ParseIndex())
    engine.Static("assets", "web/dist/assets")
    g := &Greyhound{
        *engine,
    }
    return g
}

func isAsync(request http.Request) (async bool, e error) {
    if request.Header == nil {
        return false, errors.New("greyhound.core: Request has no headers")
    }
    if request.Header.Get("X-Greyhound-Async") == "true" {
        return true, nil
    }
    return false, nil
}

func bail(c *gin.Context, e error) {
    c.AbortWithError(http.StatusInternalServerError, e)
}

func dispatch(controller controllers.Controller) gin.HandlerFunc {
    return func(context *gin.Context) {
        component, props := controller(context)
        async, err := isAsync(*context.Request)
        if err != nil {
            bail(context, err)
        }
        propsJson, err := json.Marshal(props)
        if err != nil {
            bail(context, err)
        }
        if async {
            context.JSON(http.StatusOK, gin.H{
                "component": component,
                "props": propsJson,
            })
        } else {
            context.HTML(http.StatusOK, "index.html", gin.H{
                "component": component,
                "props": propsJson,
            })
        }
    }
}

func (g *Greyhound) Register(route routes.Route) {
    g.engine.Handle(string(route.Method), route.Path, dispatch(route.Controller))
}

