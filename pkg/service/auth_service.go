package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/KoteiIto/wire-sample/pkg/repository"
)

type AuthenticateInput struct {
	Email    string
	Password string
}

type AuthenticateOutput struct {
	UserId       int64
	OnetimeToken string
}

type IAuthService interface {
	Authenticate(ctx context.Context, input *AuthenticateInput) (*AuthenticateOutput, error)
}

type AuthService struct {
	userRepository         repository.IUserRepository
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
	user, err := s.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if input.Password != user.Password {
		return nil, errors.New("passwordが一致しません")
	}

	token, err := s.onetimeTokenRepository.Issue(user)
	if err != nil {
		return nil, err
	}

	return &AuthenticateOutput{
		UserId:       user.UserId,
		OnetimeToken: token,
	}, nil
}
