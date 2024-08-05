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

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type ForgetPasswordResponse struct {
	ExpiresAt string `json:"expiresAt"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" validate:"required"`
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type ApproveUserRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	LocationId  string `json:"locationId" validate:"required"`
}

type TBXRegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Email       string `json:"email" validate:"email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Role        string `json:"role" validate:"required"`
}
