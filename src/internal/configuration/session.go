package configuration

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	SessionKeyUser = "user"
)

type SessionConfiguration struct {
	Secret string `yaml:"secret" json:"secret"`
}

func NewSessionConfiguration(secret string) *SessionConfiguration {
	return &SessionConfiguration{
		Secret: secret,
	}
}

func GetSessionConfiguration() *SessionConfiguration {
	// Might use ozzo-validation here, instead of manually checking each parameter
	secret := os.Getenv("SESSION_SECRET")

	if secret == "" {
		secret = "secret"
	}

	return NewSessionConfiguration(secret)
}

func IsAuthenticated(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	user := session.Get(SessionKeyUser)

	if user == nil {
		return false
	}

	return true
}
