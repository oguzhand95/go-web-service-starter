package controller

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/oguzhand95/go-web-service-starter/src/internal/configuration"
	"net/http"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) GetHome(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "home", gin.H{
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})
}
