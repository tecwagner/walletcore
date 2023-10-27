package authenticationUser

type AuthenticationInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationOutputDTO struct {
	Token string `json:"token"`
}
