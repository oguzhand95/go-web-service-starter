package configuration

import (
	"flag"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	SessionKeyUser = "user"
)

type SessionConfiguration struct {
	Secret *string `yaml:"secret" json:"secret"`
}

func (sc SessionConfiguration) Validate() error {
	return validation.ValidateStruct(&sc,
			validation.Field(&sc.Secret, validation.NilOrNotEmpty),
		)
}

func NewSessionConfiguration(secret *string) *SessionConfiguration {
	return &SessionConfiguration{
		Secret: secret,
	}
}

func GetSessionConfiguration() *SessionConfiguration {
	sessionSecret := flag.String("session-secret", "secret", "session secret")

	return NewSessionConfiguration(sessionSecret)
}

func IsAuthenticated(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	user := session.Get(SessionKeyUser)

	if user == nil {
		return false
	}

	return true
}
