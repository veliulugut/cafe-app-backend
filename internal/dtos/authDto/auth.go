package authDto

type LoginRequest struct {
	UniqueIdentifier string `json:"uniqueIdentifier" validate:"required"`
	Password         string `json:"password" validate:"required"`
}

type LoginResponse struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Token       string `json:"token"`
}
