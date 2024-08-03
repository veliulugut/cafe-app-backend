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

type KeycloakAttributes struct {
	CompanyId            []string `json:"companyId"`
	BranchId             []string `json:"branchId"`
	IdentityNumber       []string `json:"identityNumber"`
	PhoneNumber          []string `json:"phoneNumber"`
	EmergencyName        []string `json:"emergencyName"`
	EmergencyPhoneNumber []string `json:"emergencyPhoneNumber"`
	IBAN                 []string `json:"iban"`
	Responsibility       []string `json:"responsibility"`
	Address              []string `json:"address"`
	BloodType            []string `json:"bloodType"`
	Other                []string `json:"other"`
	Roles                []string `json:"roles"`
	Permissions          []string `json:"permissions"`
}
