package request

type LoginRequest struct {
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

type RegisterRequest struct {
	Email           string `json:"email" xml:"email"`
	Password        string `json:"password" xml:"password"`
	ConfirmPassword string `json:"confirmPassword" xml:"confirmPassword"`
}
