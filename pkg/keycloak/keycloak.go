package keycloak

import (
	"cafe-app-backend/internal/dtos/authDto"
	"context"
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

var _ Keycloak = (*KeycloakService)(nil)

func New() *KeycloakService {
	return &KeycloakService{
		baseurl:             os.Getenv("KEYCLOAK_BASE_URL"),
		realm:               os.Getenv("KEYCLOAK_REALM"),
		restApiClientId:     os.Getenv("KEYCLOAK_REST_API_CLIENT_ID"),
		restApiClientSecret: os.Getenv("KEYCLOAK_REST_API_CLIENT_SECRET"),
	}
}

type KeycloakService struct {
	baseurl             string
	realm               string
	restApiClientId     string
	restApiClientSecret string
}

// GetCompanyUsers implements Keycloak.
func (k *KeycloakService) GetCompanyUsers(ctx context.Context, companyId int) ([]*gocloak.User, error) {
	panic("unimplemented")
}

// GetUsers implements Keycloak.
func (k *KeycloakService) GetUsers(ctx context.Context, params gocloak.GetUsersParams) ([]*gocloak.User, error) {
	panic("unimplemented")
}

func (k *KeycloakService) LoginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(k.baseurl)
	token, err := client.LoginClient(ctx, k.restApiClientId, k.restApiClientSecret, k.realm)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return token, nil
}

func (k *KeycloakService) DeleteUseR(ctx context.Context, userId string) error {
	panic("unimplemented")
}

func (k *KeycloakService) Login(ctx context.Context, username string, password string) (*gocloak.JWT, error) {
	client := gocloak.NewClient(k.baseurl)

	token, err := client.Login(ctx, k.restApiClientId, k.restApiClientSecret, k.realm, username, password)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return token, nil
}

func (k *KeycloakService) Register(ctx context.Context, user authDto.KeycloakRegister) (*gocloak.User, error) {
	token, err := k.LoginRestApiClient(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	client := gocloak.NewClient(k.baseurl)

	userId, err := client.CreateUser(ctx, token.AccessToken, k.realm, user.User)
	if err != nil {
		log.Error(err)
		return nil, errors.Wrap(err, "unable to create the user")
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, k.realm, user.Password, false)
	if err != nil {
		log.Error(err)
		return nil, errors.New("unable to set the pasword for the user")
	}

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, k.realm, userId)
	if err != nil {
		log.Error(err)
		return nil, errors.New("unable to get recently created user")
	}

	return userKeycloak, nil
}

func (k *KeycloakService) UpdateUser(ctx context.Context, user *gocloak.User) error {
	token, err := k.LoginRestApiClient(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	client := gocloak.NewClient(k.baseurl)

	err = client.UpdateUser(ctx, token.AccessToken, k.realm, *user)
	if err != nil {
		return fmt.Errorf("unable to update user: %v", err)
	}

	return nil
}

func (k *KeycloakService) UpdateUserPassword(ctx context.Context, user *gocloak.User, password string) error {
	token, err := k.LoginRestApiClient(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	client := gocloak.NewClient(k.baseurl)

	err = client.SetPassword(ctx, token.AccessToken, *user.ID, k.realm, password, false)
	if err != nil {
		log.Error(err)
		return errors.New("unable to set the pasword for the user")
	}

	return nil

}

func (k *KeycloakService) ApproveUser(ctx context.Context, userId string) error {
	token, err := k.LoginRestApiClient(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	client := gocloak.NewClient(k.baseurl)

	err = client.UpdateUser(ctx, token.AccessToken, k.realm, gocloak.User{
		ID:      &userId,
		Enabled: gocloak.BoolP(true),
	})

	if err != nil {
		log.Error(err)
		return errors.Wrap(err, "unable to update user")
	}

	return nil
}
