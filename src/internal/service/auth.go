package service

import (
	"errors"
	"github.com/oguzhand95/go-web-service-starter/src/internal/model"
	"github.com/oguzhand95/go-web-service-starter/src/internal/model/request"
	"github.com/oguzhand95/go-web-service-starter/src/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (as *AuthService) Login(loginRequest *request.LoginRequest) error {
	if strings.Trim(loginRequest.Email, " ") == "" || strings.Trim(loginRequest.Password, " ") == "" {
		return errors.New("kullanıcı adı veya parola boş olamaz")
	}

	user, err := as.UserRepository.GetByFieldMail(loginRequest.Email)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		return err
	}

	if loginRequest.Email == user.Mail {
		return nil
	}

	return errors.New("giris sırasında hata meydana geldi")
}

func (as *AuthService) Register(registerRequest *request.RegisterRequest) error {
	if strings.Trim(registerRequest.Email, " ") == "" ||
		strings.Trim(registerRequest.Password, " ") == "" ||
		strings.Trim(registerRequest.ConfirmPassword, " ") == "" {
		return errors.New("Gerekli tüm alanları doldurunuz")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	err = as.UserRepository.Register(&model.User{
		Mail:     registerRequest.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return errors.New("user couldn't be saved into database")
	}

	return nil
}
