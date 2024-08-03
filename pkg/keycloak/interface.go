package keycloak

import (
	"cafe-app-backend/internal/dtos/authDto"
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type Keycloak interface {
	Register(ctx context.Context, user authDto.KeycloakRegister) (*gocloak.User, error)
	Login(ctx context.Context, username string, password string) (*gocloak.JWT, error)
	UpdateUser(ctx context.Context, user *gocloak.User) error
	UpdateUserPassword(ctx context.Context, user *gocloak.User, password string) error
	DeleteUseR(ctx context.Context, userId string) error
	LoginRestApiClient(ctx context.Context) (*gocloak.JWT, error)
	ApproveUser(ctx context.Context, userId string) error
	GetUsers(ctx context.Context, params gocloak.GetUsersParams) ([]*gocloak.User, error)
	GetCompanyUsers(ctx context.Context, companyId int) ([]*gocloak.User, error)
	ResetPassword(ctx context.Context, user authDto.KeycloakUpdateUserPassword) error
	FindUseR(ctx context.Context, username string) (*gocloak.User, error)
	UpdateUserWithoutEnable(ctx context.Context, user *gocloak.User) error
	RemoveRoleFromUser(ctx context.Context, userId, roleName string) error
	AssignNewRoleToUser(ctx context.Context, userId, roleName string) error
	GetBrachUseR(ctx context.Context, branchId int) ([]*gocloak.User, error)
}
