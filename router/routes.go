package router

import (
	"net/http"
	"github.com/vashu992/Google-Translator/constant"
	"github.com/vashu992/Google-Translator/controller"
)

// this routes represent user
var translateRoutes = Routes{
	Route{"Translate from to language", http.MethodPost, constant.TranslatePath, controller.Translate},
}