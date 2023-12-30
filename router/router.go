package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	method      string
	Path        string // localhost:8000/helloword
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) Translator(g *gin.RouterGroup) {
	group := g.Group("/" + os.Getenv("TranslateRoute"))
	for _, singleRoute := range translateRoutes {
		switch singleRoute.method {
		case http.MethodPost:
			group.POST(singleRoute.Path, singleRoute.HandlerFunc)
		case http.MethodGet:
			group.GET(singleRoute.Path, singleRoute.HandlerFunc)
		case http.MethodPut:
			group.PUT(singleRoute.Path, singleRoute.HandlerFunc)
		case http.MethodDelete:
			group.DELETE(singleRoute.Path, singleRoute.HandlerFunc)
		}
	}
}

func Routing() {
	r := routes{
		router: gin.Default(),
	}
	grouping := r.router.Group(os.Getenv("ApiVersion"))
	r.Translator(grouping)
	r.router.Run(":" + os.Getenv("port"))
}
