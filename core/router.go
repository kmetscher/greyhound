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

func Router() *Greyhound {
    engine := gin.Default()
    // TODO: extract indexing and static asset loading to end user bootstrapping
    // possibly also just tie to /dist after build; am I all in on Vue or want to be agnostic?
    // HACK: just declare those paths as constructor params?
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

// TODO: shit myself and die with elegance and grace
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
                "props": string(propsJson),
            })
        }
    }
}

func (g *Greyhound) Register(route routes.Route) {
    g.engine.Handle(string(route.Method), route.Path, dispatch(route.Controller))
}

func (g *Greyhound) Run(address string) {
    g.engine.Run(address)
}

