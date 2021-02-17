package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhand95/go-web-service-starter/src/internal/configuration"
	"net/http"
)

func AuthRequired(ctx *gin.Context) {
	if !configuration.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	ctx.Next()
}

func ForceNoAuthRequired(ctx *gin.Context) {
	if configuration.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	ctx.Next()
}
