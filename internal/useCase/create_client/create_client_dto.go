package createClient

import "time"

type CreateClientInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateClientOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
