package service

import (
	"context"
	"fmt"

	"github.com/KoteiIto/wire-sample/pkg/repository"
)

type AuthenticateInput struct {
	Email string
	Password string
}

type AuthenticateOutput struct {
	UserId int64
	OnetimeToken string
}

type IAuthService interface {
	Authenticate(ctx context.Context, input *AuthenticateInput) (*AuthenticateOutput, error)
}

type AuthService struct {
	userRepository repository.IUserRepository
	onetimeTokenRepository repository.IOnetimeTokenRepository
}

var _ IAuthService = (*AuthService)(nil)

func NewAuthService(
	userRepository *repository.UserRepository,
	onetimeTokenRepository *repository.OnetimeTokenRepository,
	) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		onetimeTokenRepository: onetimeTokenRepository,
	}
}

func NewAuthServiceFromInterface(
	userRepository repository.IUserRepository,
	onetimeTokenRepository repository.IOnetimeTokenRepository,
) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		onetimeTokenRepository: onetimeTokenRepository,
	}
}

func NewAuthServiceWithCleanup(
	userRepository *repository.UserRepository,
	onetimeTokenRepository *repository.OnetimeTokenRepository,
) (*AuthService, func()) {
	cleanup := func() {
		fmt.Println("clean up end")
	}
	return &AuthService{
		userRepository:         userRepository,
		onetimeTokenRepository: onetimeTokenRepository,
	}, cleanup
}

func (s *AuthService) Authenticate(ctx context.Context, input *AuthenticateInput) (*AuthenticateOutput, error) {

	return &AuthenticateOutput{
		UserId:       0,
		OnetimeToken: "",
	}, nil
}



