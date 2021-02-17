package controller

import (
	"fmt"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oguzhand95/go-web-service-starter/src/internal/configuration"
	"github.com/oguzhand95/go-web-service-starter/src/internal/model/request"
	"github.com/oguzhand95/go-web-service-starter/src/internal/service"
	"log"
	"net/http"
	"strconv"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) GetLogin(ctx *gin.Context) {
	var errorBool bool

	hasError, ok := ctx.GetQuery("hasError")

	if !ok {
		errorBool = false
	} else {
		var err error
		errorBool, err = strconv.ParseBool(hasError)

		if err != nil {
			errorBool = false
		}
	}

	errorMessage, ok := ctx.GetQuery("errorMessage")

	if !ok {
		errorMessage = "Hata Oluştu!"
	}

	ginview.HTML(ctx, http.StatusOK, "login", gin.H{
		"title":           "Login",
		"hasError":        errorBool,
		"errorMessage":    errorMessage,
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})

}

func (ac *AuthController) PostLogin(ctx *gin.Context) {
	loginRequest := &request.LoginRequest{}
	loginRequest.Email = ctx.PostForm("email")
	loginRequest.Password = ctx.PostForm("password")

	err := ac.authService.Login(loginRequest)

	if err != nil {
		log.Printf("Logging in failed with: %s", err.Error())
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/login?hasError=true&errorMessage=%s", "Login failed"))
		return
	}

	session := sessions.Default(ctx)
	session.Set(configuration.SessionKeyUser, loginRequest.Email)

	if err := session.Save(); err != nil {
		log.Printf("Logging in failed with: %s", err.Error())
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/login?hasError=true&errorMessage=%s", "Login failed"))
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (ac *AuthController) GetRegister(ctx *gin.Context) {
	var errorBool bool

	hasError, ok := ctx.GetQuery("hasError")

	if !ok {
		errorBool = false
	} else {
		var err error
		errorBool, err = strconv.ParseBool(hasError)

		if err != nil {
			errorBool = false
		}
	}

	errorMessage, ok := ctx.GetQuery("errorMessage")

	if !ok {
		errorMessage = "Error occurred"
	}

	ginview.HTML(ctx, http.StatusOK, "register", gin.H{
		"title":           "Register",
		"hasError":        errorBool,
		"errorMessage":    errorMessage,
		"isAuthenticated": configuration.IsAuthenticated(ctx),
	})

}

func (ac *AuthController) PostRegister(ctx *gin.Context) {
	registerRequest := &request.RegisterRequest{}
	registerRequest.Email = ctx.PostForm("email")
	registerRequest.Password = ctx.PostForm("password")
	registerRequest.ConfirmPassword = ctx.PostForm("confirmPassword")

	err := ac.authService.Register(registerRequest)

	if err != nil {
		ctx.Redirect(http.StatusFound, "/register?hasError=true&errorMessage="+err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}

func (ac *AuthController) GetLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Set(configuration.SessionKeyUser, "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1}) // this sets the cookie with a MaxAge of 0
	err := session.Save()

	if err != nil {
		log.Printf("Critical error occured during logout, please check;\n%s", err.Error())
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
