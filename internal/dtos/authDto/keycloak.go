package authDto

import "github.com/Nerzal/gocloak/v13"

type KeycloakRegister struct {
	Roles    []string `json:"roles"`
	Password string   `json:"password"`
	User     gocloak.User
}

type KeycloakUpdateUserPassword struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
