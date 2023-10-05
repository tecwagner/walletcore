package createAccount

import "time"

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID        string    `json:"id_account"`
	ClientID  string    `json:"client_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
